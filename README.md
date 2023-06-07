# Tiny Pong
Tiny Pong for micro:bit-v2. There is a [code walkthrough](https://shannondybvig.com/posts/tiny-pong-on-microbit-v2/) on my website.

![tinypong-preview](https://user-images.githubusercontent.com/69212809/231538507-e296ec65-4b80-40b6-9a70-be1418902aba.gif)

## Flash Tiny Pong to the micro:bit-v2

To flash Tiny Tetris to your micro:bit-v2, type the following from inside the directory:

```
tinygo build -o=/run/media/sdybvig/MICROBIT/flash.hex -target=microbit-v2 main.go
```

## Status

Currently the game is a human player against a naive computer opponent.
- Human player starts with the ball, whoever scores a point gets to launch the ball next
- Computer paddle bounds from side to side without consideration for ball location
- Game ends when either player accumulates 5 points, and point counts are displayed one row out from each player's paddle

## Possible Enhancements

- Computer strategy (follow ball trajectory, anticipate human location when launching ball)
- Allow human player to change direction with button input
- Mode for 2 human players (with menu to select 1 or 2 human players)
- Increase update speed over time
- Opening splash animation
- Option to play again? (maintain win count across games?)
