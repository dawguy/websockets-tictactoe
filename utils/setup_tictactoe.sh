tmux new-session -d -s tictactoe
tmux split-window -v
tmux split-window -h

# Setup my most common development setup for tictactoe
tmux select-pane -t 1
tmux send-keys 'cd ~/Documents/tictactoe;'
tmux send-keys Enter
tmux select-pane -t 2
tmux send-keys 'cd ~/Documents/tictactoe;'
tmux send-keys Enter
tmux select-pane -t 0
tmux send-keys 'cd ~/Documents/tictactoe;'
tmux send-keys Enter

tmux attach-session -t tictactoe
