SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"

mkdir -p $SCRIPT_DIR/{config,data,logs}

podman run --detach \
  --hostname localhost \
  --publish 9443:443 --publish 9080:80 --publish 9022:22 \
  --name gitlab \
  --volume $SCRIPT_DIR/config:/etc/gitlab:Z \
  --volume $SCRIPT_DIR/logs:/var/log/gitlab:Z \
  --volume $SCRIPT_DIR/data:/var/opt/gitlab:Z \
  gitlab/gitlab-ee:latest

while [ ! -f "$SCRIPT_DIR/config/initial_root_password" ]; do sleep 1; done

# wait and then enter with root and this password
grep 'Password:' $SCRIPT_DIR/config/initial_root_password
