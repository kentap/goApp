import streamlit as st
import requests
import pandas as pd

# Streamlitのタイトル
st.title("YouTube分析アプリ")

# 検索ワードの入力
search_word = st.text_input("Search for a keyword:")

# 実行回数の選択（最大20回分）
nums = st.selectbox("How many results to fetch?", list(range(1, 21)), index=0)

# APIエンドポイント（Goバックエンド）
API_URL = "http://localhost:8080/get-youtube-data"

# API呼び出しを行う関数
def fetch_youtube_data(query, nums):
    # APIに送るパラメータ
    params = {"query": query, "nums": str(nums)}
    
    try:
        # GETリクエストを送信
        response = requests.get(API_URL, params=params)
        response.raise_for_status()
        return response.json()
    except requests.exceptions.RequestException as e:
        st.error(f"Error fetching data: {e}")
        return None

# 「実行」ボタンをクリックした時の動作
if st.button("Run"):
    if search_word:
        # YouTubeデータをAPIから取得
        data = fetch_youtube_data(search_word, nums)

        # 取得できた場合、DataFrameに変換して表示
        if data:
            df = pd.DataFrame(data)
            st.write(df)
    else:
        st.warning("Please enter a search keyword.")
