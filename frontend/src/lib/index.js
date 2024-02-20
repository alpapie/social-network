// place files you want to import through the `$lib` alias in this folder.
import { log } from "console";
import { writeFileSync } from "fs";


export const  generateRandom = () =>
  Math.random().toString(36).substring(2, 15) +
  Math.random().toString(23).substring(2, 5);

  export const saveImage = async (image) => {
    console.log("here is the image", image);
    if (image.name == 'undefined' || image.size ===  0) return "";
    let extension = image.name.split(".").pop().toLowerCase();
    if (!['jpeg', 'jpg', 'png', 'gif'].includes(extension)) {
        console.error('Invalid image format. Only JPEG, PNG, and GIF are allowed.');
        return "";
    }
    const maxSize =  20 *  1024 *  1024; //  20MB in bytes
    if (image.size > maxSize) {
        console.error('Image size exceeds  20MB limit.');
        return "";
    }
    let path = `./static/uploads/${generateRandom()}.${extension}`;
    writeFileSync(path, Buffer.from(await image.arrayBuffer()));
    return path;
}