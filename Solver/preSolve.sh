cat ../Sources/finalComponentsList > ../Sources/finalNutrientsMatrixList
cat ../Sources/desiredNutrients >> ../Sources/finalNutrientsMatrixList
go run ./transpose.go ../Sources/finalNutrientsMatrixList > ./InitialNutrientMatrix
#now that we have a a nutrient matrix, we need to trim the rows that we don't want:
cat ./InitialNutrientMatrix | grep -v '$NA' > NutrientMatrix
rm InitialNutrientMatrix
go run ./generateLPSolveFile.go > lpfile
