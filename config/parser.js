// parser.js
import { parse } from 'json5';
import { promisify } from 'util';

const parseAsync = promisify(parse);

export async function parseFile(filePath) {
  try {
    const fileContent = await promisify(require('fs').readFile)(filePath, 'utf8');
    return parseAsync(fileContent);
  } catch (error) {
    if (error.code === 'ENOENT') {
      throw new Error(`File not found: ${filePath}`);
    }
    throw error;
  }
}