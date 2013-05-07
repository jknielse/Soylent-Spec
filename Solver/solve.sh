./preSolve.sh
lp_solve lpfile | grep x > temp1
cat temp1 | sed 's/.* //g' > temp2
cat temp2 | sed ':a;N;$!ba;s/\n/ /g' > temp3
NUM=0
for val in `cat temp3`
do
    if [ $val != 0 ]
    then
        echo -n "$val times "
        ./whichFood.sh $NUM
        echo
    fi
    NUM=`expr 1 + $NUM`
done
