go run ./transpose.go ../Sources/desiredNutrientsInputFile > desiredNutrientsTemp

cat ../Sources/finalComponentsList > nutrientsMatrixListLower
awk 'NR==1' desiredNutrientsTemp >> nutrientsMatrixListLower
go run ./transpose.go nutrientsMatrixListLower > nutrientsMatrixListLowerTranspose

cat ../Sources/finalComponentsList > nutrientsMatrixListUpper
awk 'NR==2' desiredNutrientsTemp >> nutrientsMatrixListUpper
go run ./transpose.go nutrientsMatrixListUpper > nutrientsMatrixListUpperTranspose


#now that we have a nutrient matrix, we need to trim the rows that we don't want:
cat ./nutrientsMatrixListUpperTranspose | grep -v '$NA' > NutrientMatrixUpper
cat ./nutrientsMatrixListLowerTranspose | grep -v '$NA' > NutrientMatrixLower
go run ./generateLPSolveFile.go > lpfiletemp
cat lpfiletemp | grep -v "NA" > lpfile

rm lpfiletemp
rm nutrientsMatrixListLower
rm nutrientsMatrixListUpper
rm nutrientsMatrixListUpperTranspose
