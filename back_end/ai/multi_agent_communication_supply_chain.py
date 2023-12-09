# =========== Copyright 2023 @ CAMEL-AI.org. All Rights Reserved. ===========
# Licensed under the Apache License, Version 2.0 (the “License”);
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an “AS IS” BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
# =========== Copyright 2023 @ CAMEL-AI.org. All Rights Reserved. ===========
import json, re

from colorama import Fore

from camel.societies import RolePlaying
from camel.utils import print_text_animated
from camel.functions.search_functions import (
    search_google,
    text_extract_from_web,
)
from camel.types import ModelType, TaskType


import asyncio
import websockets
import json

async def echo(websocket, path):
    path = "message1"
    async for message in websocket:
        try:
            data = json.loads(message)
            # Assuming data is in the format: {"SpeakerID": "...", "ReceiverID": "...", "text": "..."}
            print(f"Received message from {data['SpeakerID']} to {data['ReceiverID']}: {data['text']}")

            # Echo back the message
            await websocket.send(message)
        except json.JSONDecodeError:
            print("Received non-JSON message")

async def main():
    async with websockets.serve(echo, "localhost", 8765):
        await asyncio.Future()  # run forever

if __name__ == "__main__":
    asyncio.run(main())




def role_playing(model_type=ModelType.GPT_3_5_TURBO_16K, chat_turn_limit=50) -> None:
    request_json = {
        "outlet_id": "1",
        "outlet_location": "Lyon",
        "central_hub_location": "Paris",
        "date": "2023-12-07T15:04:05Z",
        "event": "Lavender Festival",
        "event_description": "The upcoming Lavender Festival in the region is expected to significantly increase the demand for local specialties. We anticipate a higher demand for Olive Oil and Baguette as tourists prefer local culinary experiences. Preparing additional stock of these items is advised to meet the increased customer flow.",
        "client_preferences": "Customers in Marseille show a strong preference for locally sourced products, with an emphasis on organic and artisanal options. Olive Oil and Baguettes are particularly popular, aligning with regional culinary traditions.",
        "weather": "rainy",
        "shop_inventory": {
            "olive_oil": {
                "productID": "Olive Oil",
                "quantity": 100,
                "replenishmentRate": 30,
                "maxStock": 500
            },
            "baguette": {
                "productID": "Baguette",
                "quantity": 200,
                "replenishmentRate": 50,
                "maxStock": 300
            },
            "manchego_Cheese": {
                "productID": "Manchego Cheese",
                "quantity": 150,
                "replenishmentRate": 40,
                "maxStock": 400
            },
            "black_tea": {
                "productID": "Black Tea",
                "quantity": 150,
                "replenishmentRate": 40,
                "maxStock": 500
            }
        }
    }

    context_text = json.dumps(request_json, indent=4)
    task_prompt = "Please configure the information of all goods in inventory according to the event description (you need to tell me the basis specifically)."
    assistant_answer_template = "When you are to write \"CAMEL_TASK_DONE\", please summarize your actions in the following format:\n"
    answer_template = "===== ANSWER TEMPLATE =====\n"
    convert_into_template = lambda x: {k: convert_into_template(v) if isinstance(v, dict) else "<BLANK>" for k, v in x.items()}
    answer_template += json.dumps(convert_into_template(request_json), indent=4)
    assistant_answer_template += answer_template

    ai_assistant_role = "Central Hub"
    ai_assistant_description = "You are a central hub."
    ai_user_role = "Outlet"
    ai_user_description = "You are an outlet."
    # You can use the following code to play the role-playing game
    sys_msg_meta_dicts = [
        dict(
            assistant_role=ai_assistant_role, user_role=ai_user_role,
            assistant_description=ai_assistant_description + "\n" +
            context_text, user_description=ai_user_description)
        for _ in range(2)
    ]
    role_play_session = RolePlaying(
        assistant_role_name="AI Assistant",
        assistant_agent_kwargs=dict(model_type=model_type),
        user_role_name="AI User",
        model_type=model_type,
        task_type=TaskType.ROLE_DESCRIPTION,
        task_prompt=task_prompt + "\n" + assistant_answer_template,
        with_task_specify=False,
        extend_sys_msg_meta_dicts=sys_msg_meta_dicts,
    )

    print(
        Fore.GREEN +
        f"AI Assistant sys message:\n{role_play_session.assistant_sys_msg}\n")
    print(Fore.BLUE +
          f"AI User sys message:\n{role_play_session.user_sys_msg}\n")

    print(Fore.YELLOW + f"Original task prompt:\n{task_prompt}\n")
    print(
        Fore.CYAN +
        f"Specified task prompt:\n{role_play_session.specified_task_prompt}\n")
    print(Fore.RED + f"Final task prompt:\n{role_play_session.task_prompt}\n")

    n = 0
    input_assistant_msg, _ = role_play_session.init_chat()
    while n < chat_turn_limit:
        n += 1
        assistant_response, user_response = role_play_session.step(
            input_assistant_msg)

        if assistant_response.terminated:
            print(Fore.GREEN +
                  ("AI Assistant terminated. Reason: "
                   f"{assistant_response.info['termination_reasons']}."))
            break
        if user_response.terminated:
            print(Fore.GREEN +
                  ("AI User terminated. "
                   f"Reason: {user_response.info['termination_reasons']}."))
            break

        print_text_animated(Fore.BLUE +
                            f"AI User:\n\n{user_response.msg.content}\n", 0.005)
        print_text_animated(Fore.GREEN + "AI Assistant:\n\n"
                            f"{assistant_response.msg.content}\n", 0.005)
        # print(Fore.BLUE + f"AI User:\n\n{user_response.msg.content}\n")
        # print(Fore.GREEN + "AI Assistant:\n\n"
        #       f"{assistant_response.msg.content}\n")
        match_and_replace = lambda a, b: {k: match_and_replace(a[k], b[k]) if isinstance(a[k], dict) else b[k] for k in a}

        # Match the values in the response with the request
        final_answer_json = None
        try:
            match = re.search(r'\{.*\}', user_response.msg.content, re.DOTALL)
            if match is not None:
                json_string = match.group()
                parsed_json = json.loads(json_string)
                final_answer_json = match_and_replace(request_json, parsed_json)
        except:
            warning = "Warning: The answer does not match the template."

        try:
            match = re.search(r'\{.*\}', assistant_response.msg.content, re.DOTALL)
            if match is not None:
                json_string = match.group()
                parsed_json = json.loads(json_string)
                final_answer_json = match_and_replace(request_json, parsed_json)
        except:
            warning = "Warning: The answer does not match the template."

        if "CAMEL_TASK_DONE" in user_response.msg.content or \
            "CAMEL_TASK_DONE" in assistant_response.msg.content:
            break

        input_assistant_msg = assistant_response.msg

    print(f"final_answer_json:\n{json.dumps(final_answer_json, indent=4)}")

# if __name__ == "__main__":
#     role_playing()
