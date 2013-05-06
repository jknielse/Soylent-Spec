VAR=`go run ./whichFood.go $1`
cat ../Sources/rawComponentsList | grep "$VAR" > temp
cat desiredNutrientsTemp >> temp
go run transpose.go temp > temp2
rm temp
cat temp2 | grep -v \$NA 
rm temp2
