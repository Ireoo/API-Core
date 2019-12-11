const fs = require('fs')
const path = require('path')

let package = JSON.parse(fs.readFileSync(path.join(process.cwd(), 'package.json')))
let version = package.version
let versions = version.split('.')
versions[versions.length - 1]++
package.version = versions.join('.')
fs.writeFileSync(path.join(process.cwd(), 'package.json'), JSON.stringify(package, null, 2))
