_tput() {
    printf '\['"$(tput $@)"'\]';
}

_PS1_GIT_BRANCH_CMD='
    _GIT_BRANCH="$(git branch --show-current 2>/dev/null)";
    _GIT_REF="$(git show -q --pretty=ref HEAD 2>/dev/null | grep -io '"'"'^[a-z0-9]*'"'"')";
    _GIT_BRANCH_OR_REF="${_GIT_BRANCH:-$_GIT_REF}";
    _GIT_BRANCH_OR_REF="${_GIT_BRANCH:-none}";
    printf "git:$_GIT_BRANCH_OR_REF";
';

_ps1() {
    printf "$(_tput setaf 4)";
    printf '$(date +%%H:%%M:%%S.%%3N-%%Z)';
    printf "$(_tput sgr0) ";

    printf "$(_tput setaf 3)$(_tput bold)";
    printf '$(whoami)@$(hostname)';
    printf "$(_tput sgr0) ";

    printf "[$(_tput smul)";
    printf '$('"$_PS1_GIT_BRANCH_CMD"')';
    printf "$(_tput sgr0)] ";

    printf "$(_tput setaf 1)$(_tput bold)";
    printf '$(dirs +0)';
    printf "$(_tput sgr0) ";

    printf "$(_tput bold)";
    printf '$(if [ "$(whoami)" == "root" ]; then echo "#"; else echo "$"; fi)';
    printf "$(_tput sgr0) ";
}

PS1="$(_ps1)";