# Copyright (C) 2024, Anderson Lizarazo Tellez.
# Based on: <https://www.reddit.com/r/zellij/comments/1dn4e7i/comment/la4kbqz>

function zellij_tab_name_update_pre --on-event fish_preexec
    if not set -q ZELLIJ
        return
    end

    zellij action rename-tab (string split ' ' $argv)[1]
end

function zellij_tab_name_update_post --on-event fish_postexec
    if not set -q ZELLIJ
        return
    end

    zellij action rename-tab Fish
end
