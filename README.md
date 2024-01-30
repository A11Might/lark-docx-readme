# lark-docx-readme

Use lark docx update github README.md, such as the current document: [lark-docx-readme](https://codingnow.feishu.cn/docx/CvlWddmBkoCr5txGvz1cWblAn7f).

Available as github action. It can automatically generate a new README.md each day. Which makes updating readme easier.

## **Usage**

**github action**

```bash
# This workflow will build a golang project
# For more information see: https://docs.github.com/en/actions/automating-builds-and-tests/building-and-testing-go

name: pares docx into readme

on:
  # run automatically every 24 hours
  schedule:
     - cron: "0 */24 * * *" 
  
  # allows to manually run the job at any time
  workflow_dispatch:
  
  # run on every push on the master branch
  push:
    branches:
    - main

jobs:
  update_readme:
    permissions: 
      contents: write
    runs-on: ubuntu-latest
    name: use docx update readme
    steps:
      - name: checkout
        uses: actions/checkout@v4
      - name: generate readme
        uses: A11Might/lark-docx-readme@main
        env:
          # need set actions secrets and variables first
          APP_ID: ${{ secrets.APP_ID }}
          APP_SECRET: ${{ secrets.APP_SECRET }}
          DOCUMENT_ID: ${{ secrets.DOCUMENT_ID }}
      - name: push changes
        run: |
          git config --global user.email "41898282+github-actions[bot]@users.noreply.github.com"
          git config --global user.name "github-actions[bot]"
          git add .
          git commit -m "This is an automated commit by GitHub Actions."
          git push
```

*Inspired by *[*jamesgeorge007/github-activity-readme*](https://github.com/jamesgeorge007/github-activity-readme)



