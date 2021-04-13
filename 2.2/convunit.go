package main

import (
    "flag"
    "errors"
    "fmt"
    "os"
    "strconv"
)

type convUnit func(num float64, from string, to string) (float64, error)

func main() {
    f := flag.NewFlagSet(os.Args[1], flag.ExitOnError)
    from := f.String("from", "", "convert from")
    to := f.String("to", "", "convert to")
    err := f.Parse(os.Args[2:])
    if from == nil || to == nil {
        fmt.Fprintln(os.Stderr, "No -from and/or -to flag")
        os.Exit(2)
    }
    if (err != nil) {
        fmt.Fprintln(os.Stderr, err)
    }

    var convFunc convUnit
    switch(os.Args[1]) {
        case "weight":
            convFunc = weightMain
        case "length":
            convFunc = lengthMain
        case "temp":
            convFunc = tempMain
        default:
            fmt.Fprintf(os.Stderr, "Invalid subcommand: %s\n", os.Args[1])
            os.Exit(2)
    }

    if f.NArg() > 0 {
        for _, nstr := range f.Args() {
            n, err := strconv.ParseFloat(nstr, 64)
            if (err != nil) {
                fmt.Fprintf(os.Stderr, "Parse failure for '%s': %v\n", nstr,
                    err)
                os.Exit(2)
            }
            n, _ = convFunc(n, *from, *to)
            fmt.Println(n)
        }
    } else {
        for n, done, err := nextFloat(os.Stdin); done == false; {
            if err != nil {
                fmt.Fprintln(os.Stderr, "Parse failure on stdin")
            }
            n, _ = convFunc(n, *from, *to)
            fmt.Println(n)
        }
    }
}

func weightMain(num float64, from string, to string) (float64, error) {
    switch(from) {
        case "kilogram": // noop
        case "gram":
            num /= 1000
        case "ton":
            num *= 1000
        case "pound":
            num /= 2.205
        default:
            return 0, errors.New("invalid from")
    }
    switch(to) {
        case "kilogram": // noop
        case "gram":
            num *= 1000
        case "ton":
            num /= 1000
        case "pound":
            num *= 2.205
        default:
            return 0, errors.New("invalid to")
    }
    return num, nil
}

func lengthMain(num float64, from string, to string) (float64, error) {
    switch(from) {
        case "meter": // noop
        case "centimeter":
            num /= 100
        case "decimeter":
            num /= 10
        case "yard":
            num /= 1.09361
        case "foot":
            num /= 3.281
        case "inch":
            num /= 39.37
        default:
            return 0, errors.New("invalid from")
    }
    switch(to) {
        case "meter": // noop
        case "centimeter":
            num *= 100
        case "decimeter":
            num *= 10
        case "yard":
            num *= 1.09361
        case "foot":
            num *= 3.281
        case "inch":
            num *= 39.37
        default:
            return 0, errors.New("invalid to")
    }
    return num, nil
}

func tempMain(num float64, from string, to string) (float64, error) {
    switch(from) {
        case "celsius": // noop
        case "kelvin":
            num *= 1000
        case "fahrenheit":
            num *= 9 / 5 + 32
        default:
            return 0, errors.New("invalid from")
    }
    switch(to) {
        case "celsius": // noop
        case "kelvin":
            num /= 1000
        case "fahrenheit":
            num = (num - 32) * 5 / 9
        default:
            return 0, errors.New("invalid to")
    }
    return num, nil
}

func nextFloat(f *os.File) (float64, bool, error) {
    var ret float64
    _, err := fmt.Scanf("%g", &ret)
    return ret, false, err
}
