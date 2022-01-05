package main

import "testing" //importing the testing package so that it can be used in this test package.

func TestPerimeter(t *testing.T) { //This is a function that is named TestPerimeter that has the parameters *testing.T package and its reference name is t.
	rectangle := Rectangle{10.0, 10.0} //var rectangle is implicitly declared as Rectangle with the float64 paramaters of 10.0, 10.0.
	got := Perimeter(rectangle)        // var got is implicitly equal to Perimeter with rectangle being passed in.
	want := 40.0                       //want is a var that is being used as the expectation for that test what it wants to see and it is = to 40.

	if got != want { //here we area starting a if statement for the error message its checking if got is not equal to want then it intiates the next line.
		t.Errorf("got %.2f want %.2f", got, want) //so if got and want are not equal then t (testting var name) calls Errorf (from the testing package) and returns a string with the %.2f operator meaning default for with 2 decimal point with no exponient followed by the two var got and want.
	}
}

func TestArea(t *testing.T) { // this is a function call testarea the names testing pack t and call a pointer to the actual testing.T package not just a copy and has no return args.

	areaTests := []struct { //areatests is struct of the the names and type of the objects. Holds the info for all of our shapes.
		name    string  // decaling name as a string.
		shape   Shape   //declaring shape as a Shape capital S cause its a xport from shapes.go Shape interface.
		hasArea float64 // declaring hasArea as a flaot64.
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0}, //the next three lines are our shapes first the name followed by a string the actual shape name, then the shape struct cap for export followed by two fields the width, height and hasarea and their float64 values.
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests { // this is the start of a four loop with one undeclared arg and the other being tt and it is going to range through the areaTest struct. ******unsure about how the tt works******.
		// using tt.name from the case to use it as the `t.Run` test name
		t.Run(tt.name, func(t *testing.T) { // this is a test run calling the testing.T package referencing areaTest struct neam field. (note next lines are within the return () not the normal {})
			got := tt.shape.Area() //got is a var that contains shape from the areatest struct and the area method from the interface.
			if got != tt.hasArea { //this is the start of a if statement saying if got is not = to hasArea then perform the next line demand. needs to have a shape name from the struct and a Area that then matches the hasArea float
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea) //this is the error that gets displyed if got and want arent the same via t.Errorf from the testing package %#V print the default vaule then %g and got and want ti increase accuracy then the actual field names got, want tt.hasarea.
			}
		})

	}

	checkArea := func(t testing.TB, shape Shape, want float64) { // so i think this is just the shittier version of the code above (non refactored?) i believe minus it doesnt have triangles addded to it test still passes without it.
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		checkArea(t, rectangle, 72)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})
}
