const reverseWord = (input) => {
  let splitInput = input.split(" ");
  let result = [];

  for (let i = 0; i < splitInput.length; i++) {
    let indexedInputSplit = splitInput[i].split("");
    let lowerIndexedInputRev = splitInput[i].toLowerCase().split("").reverse();

    if (indexedInputSplit[0] == indexedInputSplit[0].toUpperCase()) {
      lowerIndexedInputRev[0] = lowerIndexedInputRev[0].toUpperCase();
    }
    result = [...result, lowerIndexedInputRev.join("")];
  }

  return result.join(" ");
};

// console.log(reverseWord("I am A Great human"));