# Load environment variables from SystemD.
# User environment variables can be set on `~/.config/environment.d`.

set -l GENERATOR /usr/lib/systemd/user-environment-generators/30-systemd-environment-d-generator

for var in ($GENERATOR)
    set -l tokens (string split --max 1 '=' "$var")

    set -l name "$tokens[1]"
    set -l value "$tokens[2]"

    set --export "$name" "$value"
end
