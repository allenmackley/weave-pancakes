const flipGroup = (pancakes, flipToIndex) => {
  const pSub = pancakes
    .slice(0, flipToIndex + 1)
    .reverse()
    .map(pancake => pancake === '-' ? '+' : '-');
  pancakes.splice(0, flipToIndex + 1, ...pSub);
};

const flipGroupsFrom = (pancakes, flipFromIndex) => {
  let numFlips = 0;
  let flipToIndex = -1;
  if (pancakes[0] === '+') {
    for (let i = 0; i < pancakes.length; i++) {
      if (pancakes[i] === '+') {
        flipToIndex = i;
      } else {
        break;
      }
    }
  }
  if (flipToIndex > -1) {
    flipGroup(pancakes, flipToIndex);
    numFlips++;
  }
  flipGroup(pancakes, flipFromIndex);
  numFlips++;
  return numFlips;
}

const makeStackHappy = (string) => {
  const pancakes = string.split('');
  const len = pancakes.length;
  let numFlips = 0;
  let pos = len - 1;
  while (pos >= 0 && numFlips <= len) {
    const needToFlip = pancakes[pos] === '-';
    if (needToFlip) {
      numFlips += flipGroupsFrom(pancakes, pos);
    }
    pos--;
  }
  return numFlips;
};

const tests = (t, testCases) => {
  if (t < 1) {
    return "Input error: There must be at least one test case."
  }
  if (t > 100) {
    return "Input error: Too many test cases. Max 100.";
  }
  const validSet = testCases.every(s => /^[\+\-]+$/.test(s));
  if (!validSet) {
    return "Input error: Test case strings must include only '+' and '-' characters.";
  }
  return testCases
    .map(s => makeStackHappy(s))
    .map((result, index) => `Case #${index+1}: ${result}`);
}

const testCases = [
  '-',
  '-+',
  '+-',
  '+++',
  '--+-',
  // '-+-+--',
  // '---++--',
];
const t = testCases.length;
console.log('results:', tests(t, testCases));