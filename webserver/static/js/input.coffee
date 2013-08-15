initializeEventHandlers = () ->
    console.log("Setting up event handlers...")
    $('#chat_send').click(handleChat)

handleChat = () ->
    username = $('#username').text()
    message = $('#chat_message').val()
    sendChat(username, message)
    $('#chat_message').val("")

MOVE_AMOUNT = 5

movePlayer :: (Num, Num) -> Any
movePlayer = (dx, dy) ->
    sendMove dx, dy

moveUp = (event) -> movePlayer 0, -MOVE_AMOUNT
moveDown = (event) -> movePlayer 0, MOVE_AMOUNT
moveLeft = (event) -> movePlayer -MOVE_AMOUNT, 0
moveRight = (event) -> movePlayer MOVE_AMOUNT, 0

# key event handler registration

kd.UP.down(moveUp)
kd.DOWN.down(moveDown)
kd.LEFT.down(moveLeft)
kd.RIGHT.down(moveRight)

@startInput = () ->
    initializeEventHandlers()
    setInterval kd.tick, 40
