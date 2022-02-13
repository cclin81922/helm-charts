# 在 macOS 安裝與設定 nodejs

 1. 安裝 nvm (用來安裝與切換 nodejs 版本)
 2. 使用 nvm 安裝 `nodejs 16.14.0` (隨附 `npm 8.3.1`)
 3. 建立專案群目錄 `nodejs-16.14.0-npm-webpack-babel-gulp-typejs-express-react`
 4. 切換工作目錄到專案群目錄
 5. 使用 nvm 指定使用 `nodejs 16.14.0` (每次都要執行)
 6. 使用 npm 建立並初始化專案目錄 `my-package1` (用 `foo` 樣板)
 7. 切換工作目錄到專案目錄
 8. 使用 npm 安裝套件 `webpack babel gulp typejs express react`
 9. 初始化 for developing express backend application
10. 初始化 for developing react frontend application

# 6. 使用 npm 建立並初始化專案目錄

```
npm init foo my-package1
```

Doc https://docs.npmjs.com/cli/v7/commands/npm-init

# 8. 使用 npm 安裝套件

## --save-dev

```
# webpack
npm install --save-dev webpack

# babel (TODO 根據後續實際開發流程需求再安裝所需的套件)
npm install --save-dev @babel/core @babel/cli @babel/preset-env

# gulp (TODO 根據後續實際開發流程需求再安裝所需的套件)
npm install --save-dev gulp gulp-cli gulp-webserver

# typejs
npm install --save-dev typescript
```

## --save

```
# express (TODO 需要再安裝 typejs 所需的套件)
npm install express

# react (TODO 需要再安裝 typejs 所需的套件)
npm install react react-dom
```

## 參考文件

* npm install babel https://babeljs.io/docs/en/usage/
* npm install gulp https://gulpjs.com/docs/en/getting-started/quick-start/

# 9. 初始化 for developing express backend application

# 10. 初始化 for developing react frontend application

# npm initializer

1. create-express-typescript-application 10 months ago  https://www.npmjs.com/package/create-express-typescript-application
2. create-react-app-ts                   3 months ago   https://www.npmjs.com/package/create-react-app-ts

Try run exporess-ts

```
npm init express-typescript-application my-package5
cd my-package5
npm run dev
curl http://localhost:8080
```

Try run react-ts

```
mkdir my-package6 && cd my-package6 && npm init react-app-ts && npm install yarn --legacy-peer-deps
cd my-package6
npm run start
curl http://localhost:8080
```
