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

from camel.configs import ChatGPTConfig, FunctionCallingConfig
from camel.societies import RolePlaying
from camel.utils import print_text_animated
from camel.functions import MATH_FUNCS, SEARCH_FUNCS
from camel.types import ModelType, TaskType


import asyncio
import websockets
import json
import logging

import asyncio
import websockets
import json
import logging

async def send_message1(websocket, message_text):
    msg = ""
    for char in message_text:
        msg += char
        print(f"Sending message: {msg}")
        message = {
            "SpeakerID": "1",
            "ReceiverID": "5",
            "text": msg
        }
        try:
            await websocket.send(json.dumps(message))
            await asyncio.sleep(0.05)
        except Exception as e:
            logging.error(e)

async def websocket_handler(websocket, path, message_text):
    await send_message1(websocket, message_text)

async def send_streaming_message(message_text):
    handler = lambda ws, path: websocket_handler(ws, path, message_text)
    async with websockets.serve(handler, "localhost", 8080):
        await asyncio.Future()  # run forever

if __name__ == "__main__":
    logging.basicConfig(level=logging.INFO)
    custom_message = "Hello, this is the message text!"
    asyncio.run(send_streaming_message(custom_message))



PATH_USER_ID_1 = "message1"
PATH_USER_ID_2 = "message2"
PATH_USER_ID_3 = "message3"
PATH_USER_ID_4 = "message4"
PATH_ASSISTANT_ID_1 = "message5"
PATH_ASSISTANT_ID_2 = "message6"
PATH_ASSISTANT_ID_3 = "message7"
PATH_ASSISTANT_ID_4 = "message8"
USER_PATHS = [PATH_USER_ID_1, PATH_USER_ID_2, PATH_USER_ID_3, PATH_USER_ID_4]
ASSISTANT_PATHS = [PATH_USER_ID_1, PATH_USER_ID_2, PATH_USER_ID_3, PATH_USER_ID_4]

def role_playing(model_type=ModelType.GPT_3_5_TURBO_16K, chat_turn_limit=50, request_json=None, central_hub_json=None) -> None:
    if request_json is None:
        request_json = {
            "outlet_id": "1",
            "outlet_location": "Lyon",
            "central_hub_location": "Paris",
            "date": "2023-12-07T15:04:05Z",
            "event": "Lavender Festival",
            "event_description": "The upcoming Lavender Festival in the region is expected to significantly increase the demand for local specialties. We anticipate a higher demand for Olive Oil and Baguette as tourists prefer local culinary experiences. Preparing additional stock of these items is advised to meet the increased customer flow.",
            "client_preferences": "Customers in Marseille show a strong preference for locally sourced products, with an emphasis on organic and artisanal options. Olive Oil and Baguettes are particularly popular, aligning with regional culinary traditions.",
            "weather": "rainy",
            "outlet_inventory": {
                "olive_oil": {
                    "current_storage_amount": 100,
                    "daily_replenishment_without_envent_from_central_hub": 30,
                    "max_warehouse_capacity": 500
                },
                "baguette": {
                    "current_storage_amount": 200,
                    "daily_replenishment_without_envent_from_central_hub": 50,
                    "max_warehouse_capacity": 300
                },
                "manchego_cheese": {
                    "current_storage_amount": 150,
                    "daily_replenishment_without_envent_from_central_hub": 40,
                    "max_warehouse_capacity": 400
                },
                "black_tea": {
                    "current_storage_amount": 150,
                    "daily_replenishment_without_envent_from_central_hub": 40,
                    "max_warehouse_capacity": 500
                }
            }
        }
    post_outlet_inventory_json = request_json['outlet_inventory']
    response_json = {
        "outlet_inventory": {
            "olive_oil": {
                "future_storage_amount": "<NUM>",
                "specific_reason_of_replenishment": "<STRING>"
            },
            "baguette": {
                "future_storage_amount": "<NUM>",
                "specific_reason_of_replenishment": "<STRING>"
            },
            "manchego_cheese": {
                "future_storage_amount": "<NUM>",
                "specific_reason_of_replenishment": "<STRING>"
            },
            "black_tea": {
                "future_storage_amount": "<NUM>",
                "specific_reason_of_replenishment": "<STRING>"
            }
        },
        "transportation_duration": "<NUM> day"
    }
    # Add the central hub inventory to the request
    input_json = request_json | central_hub_json
    print(Fore.RED + f"input_json:\n{json.dumps(input_json, indent=4)}\n")


    context_text = "===== CONTEXT =====\n" + json.dumps(input_json, indent=4) + """
The \"historical_daily_replenishment_amount_from_central_hub\" means the average daily replenishment amount from the central hub to the outlet in the past. So it could be used as a reference for the replenishment amount in the future.
The \"max_warehouse_capacity\" means the maximum capacity of the warehouse of the outlet.
The \"specific_reason_of_replenishment\" means the specific reason of replenishment for the outlet (the decisions made by the central hub) at present.
THe current storage amount of the outlet should be less than the maximum capacity of the warehouse of the outlet.
While making decisions, the central hub should first consider the neccessary information in the context, and then predict what is the unknown demand of outlet in the event.
"""
    task_prompt = "In order to help the outlet to handle the upcoming events well," + \
        " please make decisions based on the known information (you need to show the basis and the thoughts specifically). And finally before the \"CAMEL_TASK_DONE\", the AI assistant (Event Logistics Coordinator of Outlet) MUST strictly adhere to the structure of the JSON TEMPLATE, ONLY fill in the BLANKs, and DO NOT alter or modify any other part of the template.\n"
    assistant_answer_template = "Even if some BLANKs/NUMs/STRINGs are not moentioned (for example, \"specific_reason_of_replenishment\": \"<BLANCK>\") in the conversation, you must fill them with sertain values or strings.\n"
    answer_template = "===== JSON TEMPLATE =====\n"
    answer_template += json.dumps(response_json, indent=4)
    assistant_answer_template += answer_template

    ai_user_role = "Inventory Management Specialist of Central Hub"
    ai_user_description = "This expert should have strong organizational skills, attention to detail, and a deep understanding of supply chain and inventory management systems. They should be proficient in inventory tracking software and have the ability to analyze stock levels to ensure availability for events. Their duties would include categorizing goods, forecasting demand based on the event description, and configuring the inventory system to reflect accurate information for event-specific requirements."
    ai_assistant_role = "Event Logistics Coordinator of Outlet"
    ai_assistant_description = "The expert must have experience in event planning and logistics, with a knack for coordinating with multiple stakeholders. They should have competencies in project management, communication, and problem-solving. Their duties involve understanding the event description to determine the necessary goods, liaising with the Inventory Management Specialist to ensure proper stock levels, and overseeing the setup to meet the event's needs."

    # You can use the following code to play the role-playing game
    function_list = [*MATH_FUNCS, *SEARCH_FUNCS]
    # assistant_model_config = FunctionCallingConfig.from_openai_function_list(
    #     function_list=function_list,
    #     kwargs=dict(temperature=0.0),
    # )
    assistant_model_config = ChatGPTConfig(temperature=0.0)
    user_model_config = ChatGPTConfig(temperature=0.0)
    sys_msg_meta_dicts = [
        dict(
            assistant_role=ai_assistant_role, user_role=ai_user_role,
            assistant_description=ai_assistant_description + "\n" +
            context_text, user_description=ai_user_description)
        for _ in range(2)
    ]
    role_play_session = RolePlaying(
        assistant_role_name=ai_assistant_role,
        user_role_name=ai_user_role,
        assistant_agent_kwargs=dict(
            model_type=model_type,
            model_config=assistant_model_config,
            function_list=function_list,
        ),
        user_agent_kwargs=dict(
            model_type=model_type,
            model_config=user_model_config,
        ),
        model_type=model_type,
        task_type=TaskType.ROLE_DESCRIPTION,
        task_prompt=task_prompt + "\n" + assistant_answer_template,
        with_task_specify=False,
        extend_sys_msg_meta_dicts=sys_msg_meta_dicts,
    )

    # print(
    #     Fore.GREEN +
    #     f"AI Assistant sys message:\n{role_play_session.assistant_sys_msg}\n")
    # print(Fore.BLUE +
    #       f"AI User sys message:\n{role_play_session.user_sys_msg}\n")

    print(Fore.YELLOW + f"Original task prompt:\n{task_prompt}\n")
    print(
        Fore.CYAN +
        f"Specified task prompt:\n{role_play_session.specified_task_prompt}\n")
    print(Fore.RED + f"Final task prompt:\n{role_play_session.task_prompt}\n")

    final_answer_json = None

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
                            f"{ai_user_role}:\n\n{user_response.msg.content}\n", 0.005)
        print_text_animated(Fore.GREEN + f"{ai_assistant_role}:\n\n"
                            f"{assistant_response.msg.content}\n", 0.005)

        match_and_replace = lambda a, b: {k: match_and_replace(a[k], b[k]) if isinstance(a[k], dict) else b[k] for k in a}

        # Match the values in the response with the request
        try:
            _final_answer_json = None
            match = re.search(r'\{.*\}', user_response.msg.content, re.DOTALL)
            if match is not None:
                json_string = match.group()
                parsed_json = json.loads(json_string)
                _final_answer_json = match_and_replace(response_json, parsed_json)
        except:
            warning = "The user response is not in the correct format. Please check the user response."
        if _final_answer_json is not None:
            final_answer_json = _final_answer_json

        try:
            _final_answer_json = None
            match = re.search(r'\{.*\}', assistant_response.msg.content, re.DOTALL)
            if match is not None:
                json_string = match.group()
                parsed_json = json.loads(json_string)
                _final_answer_json = match_and_replace(response_json, parsed_json)
        except:
            warning = "The assistant response is not in the correct format. Please check the assistant response."
        if _final_answer_json is not None:
            final_answer_json = _final_answer_json

        if "CAMEL_TASK_DONE" in user_response.msg.content or \
            "CAMEL_TASK_DONE" in assistant_response.msg.content:
            break

        input_assistant_msg = assistant_response.msg

    # Convert string into JSON

    outlet_inventory_json = final_answer_json['outlet_inventory']

    # Calculate the changed replenishment amount from central hub
    changed_replenishment_amount_from_central_hub = {}
    for item in outlet_inventory_json:
        if item in post_outlet_inventory_json:
            # Subtract the values of the same items
            changed_replenishment_amount_from_central_hub[item] = {
                "changed_replenishment_amount_from_central_hub":
                int(outlet_inventory_json[item]["future_storage_amount"]) -
                int(post_outlet_inventory_json[item]["current_storage_amount"])
            }
            # Add the values of the same items
            central_hub_json["central_hub_inventory"][item]["current_storage_amount"] -= \
                changed_replenishment_amount_from_central_hub[item]["changed_replenishment_amount_from_central_hub"]

    duration_str = final_answer_json.get("transportation_duration", "1 day")  # Default to '1 day' if not found
    # Find all numbers in the string (assuming the first one is the day count)
    numbers = [int(s) for s in duration_str.split() if s.isdigit()]
    transportation_duration_json = {}
    transportation_duration_json["transportation_duration"] = numbers[0] if len(numbers) > 0 else 1

    final_answer_json = {
        "outlet_inventory": changed_replenishment_amount_from_central_hub,
        "central_hub_inventory": central_hub_json,
        "transportation_duration": transportation_duration_json["transportation_duration"]
    }

    print(Fore.RED + f"final_answer_json:\n{json.dumps(final_answer_json, indent=4)}\n")
    return final_answer_json, central_hub_json

# if __name__ == "__main__":
#     cantral_hub_json = {
#         "central_hub_inventory": {
#             "olive_oil": {
#                 "current_storage_amount": 1000
#             },
#             "baguette": {
#                 "current_storage_amount": 1000
#             },
#             "manchego_cheese": {
#                 "current_storage_amount": 1000
#             },
#             "black_tea": {
#                 "current_storage_amount": 1000
#             }
#         }
#     }
#     role_playing(central_hub_json=cantral_hub_json)
