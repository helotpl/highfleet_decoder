# highfleet_decoder

Example usage:

Encrypted message:
```
% cat test1.txt 
D90KM6Z=
S7MFSSM A1M2V 9SI ZCC5M T0K50Y _________ 67 SNI1 OIFY0 I9MV6CY ___________ =AUUC6Y
```

Find cipher keys:
```
 % go run highfleet_decoder.go -find test1.txt
[18 24 28 12]
```

Decode message:
```
% go run highfleet_decoder.go -d1 18 -d2 24 -d3 28 -d4 12 test1.txt
VXSW4UR=
AVERAGE SPEED 140 ROUTE BOCHIM _________ UZ ABAD CARGO URANIUM ___________ =SIMOOM
```
