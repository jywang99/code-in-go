package main

import (
	"code-in-go/src/buildtags"
	"code-in-go/src/dts"
	fl "code-in-go/src/flows" // aliasing imports
	"code-in-go/src/funcs"
	"code-in-go/src/strs"
)

func testStrings() {
    printHeader("strs.HelloInput()", 2)
	strs.HelloInput()
    printHeader("strs.Utf8chars()", 2)
	strs.Utf8chars()
    printHeader("strs.StringsLib()", 2)
    strs.StringsLib()
}

func testDatatypes() {
    printHeader("dts.PrintCollections()", 2)
	dts.PrintCollections()
    printHeader("dts.MultipleDeclare()", 2)
	dts.MultipleDeclare()
    printHeader("dts.ConstVars()", 2)
	dts.ConstVars()
    printHeader("dts.NumberToString()", 2)
    dts.NumberToString()
    printHeader("dts.StringToBytes()", 2)
    dts.StringToBytes()
    printHeader("dts.FloatAddition()", 2)
    dts.FloatAddition()
    printHeader("dts.Division()", 2)
    dts.Division()
    printHeader("dts.IterateMaps()", 2)
    dts.IterateMaps()
    printHeader("dts.ArraysAndSlices()", 2)
    dts.ArraysAndSlices()
    printHeader("dts.PrintPointers()", 2)
    dts.PrintPointers()
    printHeader("dts.TestChangeStructFn()", 2)
    dts.TestChangeStructFn()
    printHeader("dts.TestMethods()", 2)
    dts.TestMethods()
    printHeader("dts.TestStringer()", 2)
    dts.TestStringer()
    printHeader("dts.TestPolymorphism()", 2)
    dts.TestPolymorphism()
}

func testFlows() {
    printHeader("fl.BulidErrors()", 2)
    fl.BuildErrors()
    printHeader("fl.CatchFnError()", 2)
    fl.CatchFnError()
    printHeader("fl.CatchCapitalizeError()", 2)
    fl.CatchCapitalizeError()
    printHeader("fl.CatchCustomError()", 2)
    fl.CatchCustomError()
    printHeader("fl.TestWrappedError()", 2)
    fl.TestWrappedError()
    printHeader("fl.RecoverFromPanic()", 2)
    fl.RecoverFromPanic()
    printHeader("fl.SwitchFallthrough()", 2)
    fl.SwitchFallthrough()
    printHeader("fl.ReadBytesLoop()", 2)
    fl.ReadBytesLoop()
    printHeader("fl.RangeClause()", 2)
    fl.RangeClause()
    printHeader("fl.TestDefer()", 2)
    fl.TestDefer()
}

func testFuncs() {
    printHeader("funcs.VariadicFunc()", 2)
    funcs.VariadicFunc()
}

func testBuildTags() {
    printHeader("funcs.PrintFeatures()", 2)
    buildtags.PrintFeatures()
}

func main() {
    printHeader("Test strings", 1)
    testStrings()
    printHeader("Test datatypes", 1)
    testDatatypes()
    printHeader("Test flows", 1)
    testFlows()
    printHeader("Test functions", 1)
    testFuncs()
    printHeader("Test buildtags", 1)
    testBuildTags()
}
