declare module '*.css' {}
declare module '*.png' {
  const filePath: string;
  export default filePath;
}

declare module '*.jpg' {
  const filePath: string;
  export default filePath;
}

declare module '*?raw' {
  const contents: string;
  export default contents;
}
