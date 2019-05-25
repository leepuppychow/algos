function processData(input) {
  const inputArr = input.split("\n");
  const dependencies = inputArr
      .slice(2, inputArr.length-1)
      .reduce((acc, d, i) => {
          acc[i+1] = d.split(" ").slice(1);
          return acc;
      }, {});
  const desiredPrograms = inputArr[inputArr.length-1]
      .split(" ")
      .reverse();
  const installSequence = [];
  
  return {
      dependencies,
      desiredPrograms,
      installSequence,
  };
}

function getDependencyList({ dependencies, desiredPrograms, installSequence }) { 
  desiredPrograms.forEach(p => {
      // console.log({ 
      //     desiredPrograms,
      //     program: p,
      // })
      if (dependencies[p].length) {
          getDependencyList({
              desiredPrograms: dependencies[p],
              dependencies,
              installSequence,
          });
      }
      if (!installSequence.includes(p)) {
          installSequence.push(p);    
      }
  });
  
  // console.log({ installSequence });
  return installSequence;
}


process.stdin.resume();
process.stdin.setEncoding("ascii");
_input = "";
process.stdin.on("data", function (input) {
  _input += input;
});

process.stdin.on("end", function () {
  const res = getDependencyList(processData(_input));
  console.log(res.length);
  console.log(res.join(" "));
});
