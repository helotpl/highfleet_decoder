# highfleet_decoder

How to install (assuming you have golang installed in system):
```
% go get github.com/helotpl/highfleet_decoder
```

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
[12 18 24 28]
D90KM6Z= S7MFSSM A1M2V 9SI ZCC5M T0K50Y _________ 67 SNI1 OIFY0 I9MV6CY ___________ =AUUC6Y
PROCYON= AVERAGE SPEED 140 ROUTE BOCHIM _________ UZ ABAD CARGO URANIUM ___________ =SIMOOM
```

Decode message:
```
% go run highfleet_decoder.go -d1 12 -d2 18 -d3 24 -d4 28 test1.txt
D90KM6Z= S7MFSSM A1M2V 9SI ZCC5M T0K50Y _________ 67 SNI1 OIFY0 I9MV6CY ___________ =AUUC6Y
PROCYON= AVERAGE SPEED 140 ROUTE BOCHIM _________ UZ ABAD CARGO URANIUM ___________ =SIMOOM
```
