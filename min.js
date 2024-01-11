const s = "abcabcbb";
const s2 = "bbbbb";

var trap = function (height) {
  let left = 0;
  let right = height.length - 1;
  let highestLeft = 0;
  let highestRight = 0;
  let allWater = 0;

  while (right > left) {
    const leftHeight = height[left];
    const rightHeight = height[right];

    highestLeft = Math.max(highestLeft, leftHeight);
    highestRight = Math.max(highestRight, rightHeight);

    if (highestLeft < highestRight) {
      left++;
      if (highestLeft > leftHeight) {
        const water = highestLeft - leftHeight;
        allWater += water;
      }
    } else {
      right--;
      if (highestRight > rightHeight) {
        const water = highestRight - rightHeight;
        allWater += water;
      }
    }
  }
  return allWater;
};

const height = [0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1];

console.log(trap(height));
console.log("split");
