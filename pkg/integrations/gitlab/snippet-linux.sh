#!/bin/bash
set -e

# Download GitLab Runner
curl -L -o /tmp/gitlab-runner "{{ .CliURL }}"
chmod +x /tmp/gitlab-runner

# Move to trusted path
sudo mv /tmp/gitlab-runner /usr/bin/gitlab-runner

# Fix SELinux context (no-op on non-SELinux systems)
sudo restorecon -v /usr/bin/gitlab-runner 2>/dev/null || true

# Enable Podman socket so the docker executor can reach it
sudo systemctl enable --now podman.socket

# Detect the host's upstream DNS servers and propagate them into every Podman
# container (including nested build containers created by `podman build`).
# Without this, inner build containers inherit a loopback stub address
# (127.0.0.53 / systemd-resolved) that is unreachable from inside a container,
# causing DNS resolution failures like "Could not resolve host: github.com".
_dns_servers=""
if command -v resolvectl &>/dev/null; then
    _dns_servers=$(resolvectl dns 2>/dev/null \
        | awk '{for(i=2;i<=NF;i++) print $i}' \
        | grep -E '^[0-9]+\.[0-9]+\.[0-9]+\.[0-9]+$' \
        | sort -u | tr '\n' ' ' | xargs)
fi
if [ -z "$_dns_servers" ] && command -v nmcli &>/dev/null; then
    _dns_servers=$(nmcli dev show 2>/dev/null \
        | awk '/IP4\.DNS/ {print $2}' \
        | tr '\n' ' ' | xargs)
fi
if [ -z "$_dns_servers" ]; then
    _dns_servers=$(awk '/^nameserver/ && $2 !~ /^127\./ && $2 != "::1" {print $2}' /etc/resolv.conf \
        | tr '\n' ' ' | xargs)
fi
if [ -n "$_dns_servers" ]; then
    _toml_list=""
    for _ip in $_dns_servers; do
        [ -n "$_toml_list" ] && _toml_list="${_toml_list}, "
        _toml_list="${_toml_list}\"${_ip}\""
    done
    sudo mkdir -p /etc/containers
    if [ ! -f /etc/containers/containers.conf ]; then
        printf '[containers]\ndns_servers = [%s]\n' "$_toml_list" \
            | sudo tee /etc/containers/containers.conf > /dev/null
    elif grep -q '^\[containers\]' /etc/containers/containers.conf; then
        # Scope the dns_servers check to the [containers] section only
        if awk '/^\[containers\]/{f=1;next} /^\[/{f=0} f && /^dns_servers/{found=1} END{exit !found}' \
                /etc/containers/containers.conf; then
            # Replace dns_servers only within [containers]
            awk -v "val=dns_servers = [${_toml_list}]" \
                '/^\[containers\]/{s=1} /^\[/ && !/^\[containers\]/{s=0}
                 s && /^dns_servers/{$0=val} 1' \
                /etc/containers/containers.conf \
                | sudo tee /etc/containers/containers.conf.tmp > /dev/null \
                && sudo mv /etc/containers/containers.conf.tmp /etc/containers/containers.conf
        else
            sudo sed -i "/^\[containers\]/a dns_servers = [${_toml_list}]" \
                /etc/containers/containers.conf
        fi
    else
        printf '\n[containers]\ndns_servers = [%s]\n' "$_toml_list" \
            | sudo tee -a /etc/containers/containers.conf > /dev/null
    fi
fi

{{- if .LogToJournald}}
# Set journald as the container log driver so CI job output is captured by the
# systemd journal and can be correlated with runner daemon logs via job_id.
sudo mkdir -p /etc/containers
if [ ! -f /etc/containers/containers.conf ]; then
    printf '[containers]\nlog_driver = "journald"\n' \
        | sudo tee /etc/containers/containers.conf > /dev/null
elif grep -q '^\[containers\]' /etc/containers/containers.conf; then
    if awk '/^\[containers\]/{f=1;next} /^\[/{f=0} f && /^log_driver/{found=1} END{exit !found}' \
            /etc/containers/containers.conf; then
        # Replace existing log_driver within [containers]
        awk '/^\[containers\]/{s=1} /^\[/ && !/^\[containers\]/{s=0}
             s && /^log_driver/{$0="log_driver = \"journald\""} 1' \
            /etc/containers/containers.conf \
            | sudo tee /etc/containers/containers.conf.tmp > /dev/null \
            && sudo mv /etc/containers/containers.conf.tmp /etc/containers/containers.conf
    else
        sudo sed -i '/^\[containers\]/a log_driver = "journald"' \
            /etc/containers/containers.conf
    fi
else
    printf '\n[containers]\nlog_driver = "journald"\n' \
        | sudo tee -a /etc/containers/containers.conf > /dev/null
fi
{{- end}}

# Register runner using docker executor backed by Podman
# --docker-privileged is required for Podman: containers need CAP_SYS_ADMIN to mount /proc
sudo gitlab-runner register \
  --non-interactive \
  --url "{{ .RepoURL }}" \
  --token "{{ .Token }}" \
  --name "{{ .Name }}" \
  --executor "docker" \
  --docker-image "fedora:latest" \
  --docker-host "unix:///run/podman/podman.sock" \
  --docker-privileged

{{- if not .Unsecure}}
# Create a dedicated system user for running CI jobs
sudo useradd --system \
  --shell /bin/bash \
  --create-home \
  --home-dir /home/gitlab-runner \
  gitlab-runner

RUNNER_USER=gitlab-runner
{{- else}}
RUNNER_USER={{ .User }}
{{- end}}

# Install and start as service
sudo gitlab-runner install --user="${RUNNER_USER}"
{{- if .Concurrent}}
sudo sed -i "s/^concurrent = .*/concurrent = {{.Concurrent}}/" /etc/gitlab-runner/config.toml
{{- end}}
sudo systemctl daemon-reload
sudo systemctl enable --now gitlab-runner
