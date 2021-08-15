function seekAndDestroy(arr) {
    return arr.filter(
        x =>
            !Object.values(arguments)
                .slice(1)
                .find(y => y === x)
    );
}
let result1 = seekAndDestroy([1, 2, 3, 1, 2, 3], 2, 3);
let result2 = seekAndDestroy([1, 2, 3, 5, 1, 2, 3], 2, 3);
let result3 = seekAndDestroy(["foo", "bar", 1], "foo", 1);
console.log(result1)
console.log(result2)
console.log(result3)