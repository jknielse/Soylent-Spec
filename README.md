Soylent-Spec
============

A tool for determining which foods to order, and in what quantity, in order to achieve an ideal diet.

Each source of food that can be used for soylent has an src file within the Sources folder. 
This src file contains all of the nutritional information about the source.

The goal of the mixture is specified in a special src file named Goal.src.

Once all of the src files have been specified, simply run calc.go, and a results.txt file will be generated.
This text file contains all of the different amounts of the different source materials that should be used 
to make soylent.
