curl 127.0.0.1:8090/game/start;
# X wins horizontally
curl -X POST 127.0.0.1:8090/game/place -d '{"x":1,"y":1}';
curl -X POST 127.0.0.1:8090/game/place -d '{"x":1,"y":0}';
curl -X POST 127.0.0.1:8090/game/place -d '{"x":0,"y":1}';
curl -X POST 127.0.0.1:8090/game/place -d '{"x":0,"y":0}';
curl -X POST 127.0.0.1:8090/game/place -d '{"x":2,"y":1}';
curl -X POST 127.0.0.1:8090/game/place -d '{"x":2,"y":2}';
curl -X POST 127.0.0.1:8090/game/place -d '{"x":2,"y":1}';

# O wins vertically
# curl -X POST 127.0.0.1:8090/game/place -d '{"x":1,"y":1}';
# curl -X POST 127.0.0.1:8090/game/place -d '{"x":0,"y":0}';
# curl -X POST 127.0.0.1:8090/game/place -d '{"x":1,"y":0}';
# curl -X POST 127.0.0.1:8090/game/place -d '{"x":0,"y":1}';
# curl -X POST 127.0.0.1:8090/game/place -d '{"x":2,"y":2}';
# curl -X POST 127.0.0.1:8090/game/place -d '{"x":0,"y":2}';

# X wins diagonally
# curl -X POST 127.0.0.1:8090/game/place -d '{"x":1,"y":1}';
# curl -X POST 127.0.0.1:8090/game/place -d '{"x":1,"y":0}';
# curl -X POST 127.0.0.1:8090/game/place -d '{"x":2,"y":2}';
# curl -X POST 127.0.0.1:8090/game/place -d '{"x":2,"y":0}';
# curl -X POST 127.0.0.1:8090/game/place -d '{"x":0,"y":0}';

# O wins diagonally
# curl -X POST 127.0.0.1:8090/game/place -d '{"x":2,"y":1}';
# curl -X POST 127.0.0.1:8090/game/place -d '{"x":1,"y":1}';
# curl -X POST 127.0.0.1:8090/game/place -d '{"x":2,"y":2}';
# curl -X POST 127.0.0.1:8090/game/place -d '{"x":2,"y":0}';
# curl -X POST 127.0.0.1:8090/game/place -d '{"x":1,"y":2}';
# curl -X POST 127.0.0.1:8090/game/place -d '{"x":0,"y":2}';

# No Winner
# curl -X POST 127.0.0.1:8090/game/place -d '{"x":1,"y":1}';
# curl -X POST 127.0.0.1:8090/game/place -d '{"x":0,"y":0}';
# curl -X POST 127.0.0.1:8090/game/place -d '{"x":1,"y":0}';
# curl -X POST 127.0.0.1:8090/game/place -d '{"x":1,"y":2}';
# curl -X POST 127.0.0.1:8090/game/place -d '{"x":0,"y":1}';
# curl -X POST 127.0.0.1:8090/game/place -d '{"x":2,"y":1}';
# curl -X POST 127.0.0.1:8090/game/place -d '{"x":2,"y":0}';
# curl -X POST 127.0.0.1:8090/game/place -d '{"x":0,"y":2}';
# curl -X POST 127.0.0.1:8090/game/place -d '{"x":2,"y":2}';
