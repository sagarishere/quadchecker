#! /bin/bash

# Compile the source code
cd quadmaker/quaditA
go build -o ./quadA quadA.go
mv quadA ../../
cd ../../

cd quadmaker/quaditB
go build -o ./quadB quadB.go
mv quadB ../../
cd ../../

cd quadmaker/quaditC
go build -o ./quadC quadC.go
mv quadC ../../
cd ../../

cd quadmaker/quaditD
go build -o ./quadD quadD.go
mv quadD ../../
cd ../../

cd quadmaker/quaditE
go build -o ./quadE quadE.go
mv quadE ../../
cd ../../

# Compile the quadchecker
go build -o ./quadchecker ./quadchecker.go
