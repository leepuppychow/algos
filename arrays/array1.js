const sumOfPairInArray = (sum, arr) => {
  return arr.reduce((acc, num, i) => {
    acc[num] = true;
    if (acc[sum - num]) {
      console.log("found a pair!", num, sum - num)
    }
    return acc;
  }, {})
}


sumOfPairInArray(10, [1,2,3,7,4,9,3]);