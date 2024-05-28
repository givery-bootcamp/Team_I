## プロジェクトのセットアップ

### 前提条件

このプロジェクトを実行するためには、以下のソフトウェアが必要です：

- Node.js (v20.12.2)
- npm

## プロジェクト構成

```
public
src
├── app
│   ├── App.scss
│   ├── App.tsx
│   ├── AppRouter.tsx
├── components
│   ├── Navigation.tsx
│   ├── PostList.tsx
├── features
│   ├── HelloWorld
│   │   └── tsx
│   └── index.ts
├── shared
│   ├── hooks
│   ├── models
│   ├── services
│   ├── store
├── main.scss
├── main.tsx
└── vite-env.d.ts
```

### ディレクトリとその説明

- `public` 
  - 静的なファイルを配置します。例：index.html, アイコン, その他のアセット。

- `src` 
  - ソースコードを格納するメインディレクトリ。 
- `src/app` 
  - アプリケーションのエントリーポイントやルーティング関連のファイルを配置します。 
    - App.scss: アプリケーション全体のスタイルを定義。 
    - App.tsx: メインのアプリケーションコンポーネント。 
    - AppRouter.tsx: アプリケーションのルーティング設定。 
- `src/components`
  - アプリケーション内で再利用されるプレゼンテーションコンポーネントを配置します。 
    - Navigation.tsx: サイドナビゲーションバーのコンポーネント。 
    - PostList.tsx: 投稿一覧を表示するコンポーネント。 
- `src/features`
  - 特定の機能に関連するコンポーネントやロジックを配置します。 
    - HelloWorld/tsx: HelloWorld 機能に関連するコンポーネント。 
    - index.ts: features モジュールのエントリーポイント。 
- `src/shared` 
  - アプリケーション全体で共有されるリソースやコンポーネントを配置します。 
    - `hooks`: 再利用可能なカスタムフック（例：useAuth.ts, useFetch.ts）。 
    - `models`: アプリケーションで使用されるデータモデルの定義（例：User.ts, Post.ts）。 
    - `services`: 外部APIとの通信やビジネスロジックを含むサービスクラス（例：apiService.ts, authService.ts）。 
    - `store`: グローバルな状態管理の設定（例：index.ts, slices/）。
- `main.scss`
  - グローバルスタイルシート。Tailwind CSS の設定などを含む。 
- `main.tsx`
  - アプリケーションのメインエントリーポイント。ReactDOM.render を使用してアプリケーションをルートにマウント。 
- `vite-env.d.ts`
  - Vite 用の TypeScript 環境ファイル。

## 使用方法

### コンポーネント
- Navigation.tsx: サイドナビゲーションバーを表示
- PostList.tsx: 投稿の一覧を表示


### 状態管理
React Hooksを使用してコンポーネントレベルで状態管理を行います。

### スタイル
Tailwind CSSをしようしてスタイリングを行います。スタイルシートは`scr/main.scss`に定義されています。