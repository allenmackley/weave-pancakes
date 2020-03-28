const flipGroup = (pancakes, pos) => {
  const pSub = pancakes
    .slice(0, pos + 1)
    .map(pancake => pancake === '-' ? '+' : '-');
  pancakes.splice(0, pSub.length, ...pSub);
}

const makeStackHappy = (string) => {
  const pancakes = string.split('');
  const len = pancakes.length;
  let numFlips = 0;
  let pos = len - 1;
  while (pos >= 0 && numFlips <= len) {
    const needToFlip = pancakes[pos] === '-';
    if (needToFlip) {
      flipGroup(pancakes, pos);
      numFlips++;
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
    .map((result, index) => `Case #${index}: ${result}`);
}

const testCases = [
  '-',
  '-+',
  '+-',
  '+++',
  '--+-',
];
const t = testCases.length;
console.log('results:', tests(t, testCases));