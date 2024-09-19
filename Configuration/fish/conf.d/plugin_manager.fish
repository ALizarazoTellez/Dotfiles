# Based on: <https://github.com/kidonng/fisher_path.fish>

set fisher_path ~/.local/share/fisher

# Bootstrap fisher.
if not test -d "$fisher_path"; and not set --query _fisher_bootstrapping
    set --export _fisher_bootstrapping

    echo Bootstrapping fisher...

    # Taken from: <https://github.com/jorgebucaran/fisher>
    curl -sL https://raw.githubusercontent.com/jorgebucaran/fisher/main/functions/fisher.fish | source
    and fisher update

    set --erase _fisher_bootstrapping
end

set --query _fisher_path_initialized && exit
set --global _fisher_path_initialized

set fish_complete_path $fish_complete_path[1] $fisher_path/completions $fish_complete_path[2..]
set fish_function_path $fish_function_path[1] $fisher_path/functions $fish_function_path[2..]

for file in $fisher_path/conf.d/*.fish
    if ! test -f (string replace --regex "^.*/" $__fish_config_dir/conf.d/ -- $file)
        and test -f $file && test -r $file
        source $file
    end
end
