# encoding=utf-8
# Time：2023/2/20 18:19
# by Yoake
import asyncio
import time
import re
import argparse

from EdgeGPT.src.EdgeGPT import Chatbot
import requests
import json


parser = argparse.ArgumentParser(description='Test for argparse')

parser.add_argument('--bridge', '-b', help='与后端通信的端口号，默认值为8080', default="8080")
args = parser.parse_args()

bridge = 8081


def listen_to_run(message):
    if message == "!start":
        return True
    elif message == "":
        return False
    else:
        return False


async def single_conversion(bot, prompt):
    print("Bot:")
    response = (await bot.ask(prompt=prompt))
    print(response)
    if response["item"]["result"]["value"] == "Success":
        post_answer(response["item"]["messages"][1]["adaptiveCards"][0]["body"][0]["text"])
    elif response["item"]["result"]["value"] == "Throttled":
        post_answer("已达到bing限制的提问次数")
    else:
        post_answer("出现未知错误")


def run():
    while True:
        if listen_to_run(get_question('')):
            bot = Chatbot(cookiePath='./cookie.json')
            print("开始与bing聊天吧！")
            old_prompt = ''
            while True:
                prompt = get_question(old_prompt)
                if prompt != old_prompt and prompt is not None and prompt != '':
                    old_prompt = prompt
                    if prompt == "!exit":
                        post_answer("结束与bing聊天")
                        break
                    elif prompt == "!reset":
                        post_answer("已重置对话")
                        bot.reset()
                    elif prompt == "你好":
                        post_answer("世界！")
                    else:
                        post_answer("bing正在生成回答...")
                        asyncio.run(single_conversion(bot, prompt))
                time.sleep(2)
            bot.close()
        time.sleep(2)


def get_question(old):
    response = requests.get(f"http://localhost:{args.bridge}/question")
    content = json.loads(response.text)
    question = content["question"]
    if question == old:
        return ''
    if content["type"] == "group":
        if f'at,qq={content["self"]}' in question:
            cq_list = re.findall("\[CQ:\S*\]", question)
            if not cq_list:
                question = question
            else:
                for i in cq_list:
                    question = question.replace(i, "")
        else:
            return ""
    else:
        cq_list = re.findall("\[CQ:\S*\]", question)
        if not cq_list:
            question = question
        else:
            for i in cq_list:
                question = question.replace(i, "")
    return question.strip()


def post_answer(answer):
    response = requests.post(f"http://localhost:{args.bridge}/answer", data=answer.encode("utf-8"))
    if response.status_code == 200:
        return "OK"


if __name__ == '__main__':
    run()
