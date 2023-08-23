echo go run . --color red "banana"
go run . --color red "banana"
echo ----------------------------------------------------------------
echo go run . --color=red "hello world"
go run . --color=red "hello world"
echo ----------------------------------------------------------------
echo go run . --color=green "1 + 1 = 2"
go run . --color=green "1 + 1 = 2"
echo ----------------------------------------------------------------
echo go run . ---color=yellow "(%&) ??"
go run . --color=yellow "(%&) ??"
echo ----------------------------------------------------------------
echo go run . --color=orange GuYs "HeY GuYs", in order to color GuYs
go run . --color=orange GuYs "HeY GuYs"
echo ----------------------------------------------------------------
echo go run . --color=blue B 'RGB()', in order to color just the B
go run . --color=blue B 'RGB()'
echo ----------------------------------------------------------------
echo go run . --color flag notations like: --color=red
go run . --color=red
echo ----------------------------------------------------------------
echo go run . --color="rgb(255, 0, 0)"
go run . --color="rgb(255, 0, 0)"
echo ----------------------------------------------------------------
echo go run . --color="hsl(0, 100%, 50%)"
go run . --color=hsl"(0, 100%, 50%)"
echo ----------------------------------------------------------------
