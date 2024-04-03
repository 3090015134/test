//@ts-ignore
import { sqlite } from "better-sqlite3"

export function aa(){
    console.log(2);
const db = new sqlite("./log.db")

const stmt = db.prepare(`
      SELECT *
      FROM game_dbs
    `)
const rows = stmt.all()
// const a = rows.map(row => new Todo(row.text, row.completed))

console.log(rows);

}
