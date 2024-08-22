function bubbleSort(num) {
    // let num = number.length
    for (let i = 0; i < num.length - 1; i++) {
        for (let j = i + 1; j < num.length; j++) {
            if (num[i].length < num[j].length) {
                n = num[i];
                num[i] = num[j];
                num[j] = n;
            }
        }
    }
    return num;
}
function main() {
    // let num = [0,4,3,1,5,8]
    str = ["hello", "w", "ii", "are", "love", "as"]
    // log.console(num)
    console.log(bubbleSort(str))
}

main()