# Copyright (C) 2024, Anderson Lizarazo Tellez.

if status is-interactive
    # A smarter `cd` command.
    zoxide init fish --cmd cd | source

    abbr --add ls -- eza
    abbr --add tree -- eza -T
    abbr --add vim -- nvim

    # Seems to be that `--staged` and `--cached` are the same.
    abbr --add gdf -- git diff --staged
    abbr --add gst -- git status --short --renames
end
