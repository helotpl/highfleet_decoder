# highfleet_decoder

Example usage:

Find cipher keys:
 % go run highfleet_decoder.go -find test1.txt
[18 24 28 12]

Decode message:
% go run highfleet_decoder.go -d1 18 -d2 24 -d3 28 -d4 12 test1.txt
VXSW4UR=
AVERAGE SPEED 140 ROUTE BOCHIM _________ UZ ABAD CARGO URANIUM ___________ =SIMOOM
