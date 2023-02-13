#!/bin/bash

rm -rf /home/developer/.gitconfig
ln -s /home/developer/host_home/.gitconfig /home/developer/.

rm -rf /home/developer/.ssh
ln -s /home/developer/host_home/.ssh /home/developer/.

rm -rf /home/developer/.gnupg
ln -s /home/developer/host_home/.gnupg /home/developer/.

echo '
source /home/developer/tagioalisi/.devcontainer/bashrc.sh' >> /home/developer/.bashrc
