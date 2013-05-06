VAR=`go run ./whichFood.go $1`
echo $VAR
head -n 1 ../Sources/rawComponentsList > temp
cat ../Sources/rawComponentsList | grep "$VAR" >> temp
cat ../Sources/desiredNutrients >> temp
go run transpose.go temp >> temp2
rm temp
cat temp2 | grep -v \$NA 
rm temp2
