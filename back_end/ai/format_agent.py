import re
from typing import Any, Dict, Optional, Union

from camel.agents import ChatAgent
from camel.messages import BaseMessage
from camel.prompts import TextPrompt
from camel.types import ModelType, RoleType


class FormatAgent(ChatAgent):
    r"""An agent that aims to generate formatted text.
    Args:
        model_type (ModelType, optional): The type of model to use for the
            agent. (default: :obj:`ModelType.GPT_3_5_TURBO`)
        model_config (Any, optional): The configuration for the model.
            (default: :obj:`None`)
    """

    def __init__(
        self,
        model_type: ModelType = ModelType.GPT_3_5_TURBO,
        model_config: Optional[Any] = None,
    ) -> None:
        system_message = BaseMessage(
            role_name="Format Agent",
            role_type=RoleType.ASSISTANT,
            meta_dict=None,
            content="You generate format text from a provided text.",
        )
        super().__init__(system_message, model_type, model_config)

    def run(
        self,
        user_role_name: str,
        assistant_role_name: str,
        chat_record: Union[str, TextPrompt],
        answer_template: Union[str, TextPrompt],
    ) -> Union[str, TextPrompt]:
        r"""Generate role names based on the input task prompt.

        Args:
            user_role_name (str): The role name of the user.
            assistant_role_name (str): The role name of the assistant.
            chat_record (Union[str, TextPrompt]): The chat record of two
                AIs.
            answer_template (Union[str, TextPrompt]): The answer template of
                the LLM.

        Returns:
            Dict[str, Dict[str, str]]: The generated text in the format of
                the answer template.
        """
        self.reset()

        format_generation_prompt = """You are a format agent. You are asked to generate the answer according to the CHAT RECORD and the ANSWER TEMPLATE. According to the chat record, you should analyze the chat record and extract the relevant information in order to fulfill the ANSWER TEMPLATE.
Your answer MUST strictly adhere to the structure of ANSWER TEMPLATE, ONLY fill in the BLANKs, and DO NOT alter or modify any other part of the template.
"""

        chat_record_prompt = TextPrompt("""===== CHAT RECORD =====
The chat record of two AIs is provided below, where one is the \"user\" {user_role_name} and the other is the \"assistant\" {assistant_role_name}.
The context of the conversation is about the anylysis and the calculation of the \"outlet_inventory\" and the \"central_hub_inventory\" which including specific products, and also the \"transportation_duration\".
{chat_record}
=====""")  # noqa: E501

        answer_template_prompt = TextPrompt("""===== ANSWER TEMPLATE =====
{answer_template}
=====""")  # noqa: E501

        generation_prompt = TextPrompt(format_generation_prompt +
                                       chat_record_prompt +
                                       answer_template_prompt)
        generation = generation_prompt.format(
            user_role_name=user_role_name,
            assistant_role_name=assistant_role_name,
            chat_record=chat_record,
            answer_template=answer_template)

        insights_generation_msg = BaseMessage.make_user_message(
            role_name="Format Agent", content=generation)

        response = self.step(input_message=insights_generation_msg)

        if response.terminated:
            raise RuntimeError("Format text generation failed. Error:\n" +
                               f"{response.info}")
        msg = response.msg  # type: BaseMessage

        if msg.content is None:
            raise RuntimeError("Format text generation failed. Error:\n" +
                               f"{response.info}")

        return msg.content

