/**
 * @param {number} degrees
 * @returns {number}
 */
export function degreesToRadians(degrees) {
  return degrees * (Math.PI / 180);
}

/**
 * @param {number} radians
 * @returns {number}
 */
export function radiansToDegrees(radians) {
  return radians * 57.2957795;
}

/**
 * @param {number} percentage
 * @returns {number}
 */
export function percentageToRadians(percentage) {
  return ((Math.PI * 2) / 100) * percentage;
}

/**
 * @param {number} min
 * @param {number} max
 * @returns {number}
 */
export function randomBetween(min, max) {
  return Math.random() * (max - min) + min;
}

/**
 * @param {number} t
 * @returns {number}
 */
export function spinEase(t) {
  return 1 + --t * t * t * t * t;
}

/**
 * @param {number} distance
 * @param {number} angle
 * @returns {number}
 */
export function lengthDirX(distance, angle) {
  return distance * Math.cos(angle);
}

/**
 * @param {number} distance
 * @param {number} angle
 * @returns {number}
 */
export function lengthDirY(distance, angle) {
  return distance * Math.sin(angle);
}

/**
 * @param {number} t
 * @returns {number}
 */
export function resetEase(t) {
  return t * (2 - t);
}

export function createRoundedImage(src, width = 64, height = 64) {
  return new Promise(resolve => {
    const canvas = document.createElement("canvas");
    const context = canvas.getContext("2d");
    const image = new Image();

    image.src = src;

    image.onload = () => {
      canvas.width = width;
      canvas.height = height;

      context.save();
      context.beginPath();
      context.arc(width * 0.5, height * 0.5, width * 0.5, 0, Math.PI * 2, true);
      context.closePath();
      context.clip();

      context.drawImage(image, 0, 0, width, height);

      context.beginPath();
      context.arc(0, 0, 2, 0, Math.PI * 2, true);
      context.clip();
      context.closePath();
      context.restore();

      resolve(canvas);
    };
  });
}
