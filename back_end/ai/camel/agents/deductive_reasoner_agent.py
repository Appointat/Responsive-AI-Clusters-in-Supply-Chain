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
import re
from typing import Any, Dict, Optional, Union

from camel.agents import ChatAgent
from camel.messages import BaseMessage
from camel.prompts import TextPrompt
from camel.types import ModelType, RoleType


class DeductiveReasonerAgent(ChatAgent):
    r"""An agent responsible for deductive reasoning. Model of deductive
    reasoning:
        - L: A ⊕ C -> q * B
        - A represents the known starting state.
        - B represents the known target state.
        - C represents the conditions required to transition from A to B.
        - Q represents the quality or effectiveness of the transition from
        A to B.
        - L represents the path or process from A to B.

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
            role_name="Insight Agent",
            role_type=RoleType.ASSISTANT,
            meta_dict=None,
            content="You assign roles based on tasks.",
        )
        super().__init__(system_message, model_type, model_config)

    def deduce_conditions_and_quality(
        self,
        starting_state: Union[str, TextPrompt],
        target_state: Union[str, TextPrompt],
        role_descriptions_dict: Optional[Dict[str, str]] = None,
    ) -> Dict[str, Dict[str, Optional[str]]]:
        r"""Derives the conditions and quality from the starting state and the
        target state based on the model of the deductive reasoning and the
        knowledge base.

        Args:
            starting_state (Union[str, TextPrompt]): The starting state of the
                task.
            target_state (Union[str, TextPrompt]): The target state of the
                task.
            role_descriptions_dict (Optional[Dict[str, str]], optional): The
                descriptions of the roles. (default: :obj:`None`)

        Returns:
            Dict[str, Dict[str, Optional[str]]]: The generated insights from
                the input context text.
        """
        self.reset()

        deduce_prompt = """You are a deductive reasoner. You are tasked to complete the TASK based on the THOUGHT OF DEDUCTIVE REASONING, the STARTING STATE A and the TARGET STATE B. You are given the CONTEXT CONTENT to help you complete the TASK.
Your answer MUST strictly adhere to the structure of ANSWER TEMPLATE, ONLY fill in the BLANKs, and DO NOT alter or modify any other part of the template

===== MODELING OF DEDUCTIVE REASONING =====
You are tasked with understanding a mathematical model based on the components ${A, B, C, Q, L}$. In this model: ``L: A ⊕ C -> q * B``.
- $A$ represents the known starting state.
- $B$ represents the known target state.
- $C$ represents the conditions required to transition from $A$ to $B$.
- $Q$ represents the quality or effectiveness of the transition from $A$ to $B$.
- $L$ represents the path or process from $A$ to $B$.

===== THOUGHT OF DEDUCTIVE REASONING =====
1. Define the Parameters of A and B:
    - Characterization: Before delving into transitions, thoroughly understand the nature and boundaries of both $A$ and $B$. This includes the type, properties, constraints, and possible interactions between the two.
    - Contrast and Compare: Highlight the similarities and differences between $A$ and $B$. This comparative analysis will give an insight into what needs changing and what remains constant.
2. Historical & Empirical Analysis:
    - Previous Transitions according to the Knowledge Base of GPT: (if applicable) Extract conditions and patterns from the historical instances where a similar transition from a state comparable to $A$ moved towards $B$.
    - Scientific Principles: (if applicable) Consider the underlying scientific principles governing or related to the states and their transition. For example, if $A$ and $B$ are physical states, laws of physics might apply.
3. Logical Deduction of Conditions ($C$):
    - Direct Path Analysis: What are the immediate and direct conditions required to move from $A$ to $B$?
    - Intermediate States: Are there states between $A$ and $B$ that must be transversed or can be used to make the transition smoother or more efficient? If yes, what is the content?
    - Constraints & Limitations: Identify potential barriers or restrictions in moving from $A$ to $B$. These can be external (e.g., environmental factors) or internal (properties of $A$ or $B$).
    - Resource and Information Analysis: What resources and information are required for the transition? This could be time, entity, factor, code language, software platform, unknowns, etc.
    - External Influences: Consider socio-economic, political, or environmental factors (if applicable) that could influence the transition conditions.
    - Creative/Heuristic Reasoning: Open your mind to multiple possible $C$'s, no matter how unconventional they might seem. Utilize analogies, metaphors, or brainstorming techniques to envision possible conditions or paths from $A$ to $B$.
    - The conditions $C$ should be multiple but in one sentence. And each condition should be concerned with one aspect/entity.
4. Entity/Label Recognition of Conditions ($C$):
    - Identify and categorize entities of Conditions ($C$) such as the names, locations, dates, specific technical terms or contextual parameters that might be associated with events, innovations post-2022.
    - The output of the entities/labels will be used as tags or labels for semantic similarity searches. The entities/labels may be the words, or phrases, each of them should contain valuable, high information entropy information, and should be independent.
    - Ensure that the identified entities are formatted in a manner suitable for database indexing and retrieval. Organize the entities into categories, and combine the category with its instance into a continuous phrase, without using colons or other separators.
    - Format these entities for database indexing: output the category rather than its instance/content into a continuous phrase. For example, instead of "Jan. 02", identify it as "Event time".
5. Quality Assessment ($Q$):
    - Efficiency: How efficient is the transition from $A$ to $B$, which measures the resources used versus the desired outcome?
    - Effectiveness: Did the transition achieve the desired outcome or was the target state achieved as intended?
    - Safety & Risks: Assess any risks associated with the transition and the measures to mitigate them.
    - Feedback Mechanisms: Incorporate feedback loops to continuously monitor and adjust the quality of transition, making it more adaptive.
6. Iterative Evaluation:
    - Test & Refine: Based on the initially deduced conditions and assessed quality, iterate the process to refine and optimize the transition. This might involve tweaking conditions, employing different paths, or changing resources.
    - Feedback Integration: Use feedback to make improvements and increase the quality of the transition.
7. Real-world scenarios often present challenges that may not be captured by models and frameworks. While using the model, maintain an adaptive mindset:
    - Scenario Exploration: Continuously imagine various possible scenarios, both positive and negative, to prepare for unexpected events.
    - Flexibility: Be prepared to modify conditions ($C$) or alter the path/process ($L$) if unforeseen challenges arise.
    - Feedback Integration: Rapidly integrate feedback from actual implementations to adjust the model's application, ensuring relevancy and effectiveness.

===== TASK =====
Given the starting state $A$ and the target state $B$, assuming that a path $L$ always exists between $A$ and $B$, how can one deduce or identify the necessary conditions $C$ and the quality $Q$ of the transition?

===== STARTING STATE $A$ =====
{starting_state}

===== TARGET STATE $B$ =====
{target_state}

{role_with_description_prompt}
===== ANSWER TEMPLATE =====
- Characterization and comparison of $A$ and $B$:\n<BLANK>
- Historical & Empirical Analysis:\n<BLANK>/None
- Logical Deduction of Conditions ($C$) (multiple conditions can be deduced):
    condition <NUM>:
        <BLANK>.
- Entity/Label Recognition of Conditions:\n[<BLANK>, <BLANK>, ...] (include square brackets)
- Quality Assessment ($Q$) (do not use symbols):
    <BLANK>.
- Iterative Evaluation:\n<BLANK>/None"""  # noqa: E501

        if role_descriptions_dict is not None:
            role_names = role_descriptions_dict.keys()
            role_with_description_prompt = \
                "===== ROLES WITH DESCRIPTION =====\n" + "\n".join(
                    f"{role_name}:\n{role_descriptions_dict[role_name]}\n"
                    for role_name in role_names) + "\n\n"
        else:
            role_with_description_prompt = ""
        deduce_prompt = TextPrompt(deduce_prompt)

        deduce = deduce_prompt.format(
            starting_state=starting_state, target_state=target_state,
            role_with_description_prompt=role_with_description_prompt)

        conditions_and_quality_generation_msg = \
            BaseMessage.make_user_message(role_name="Deductive Reasoner",
                                          content=deduce)

        response = self.step(
            input_message=conditions_and_quality_generation_msg)

        if response.terminated:
            raise RuntimeError("Deduction failed. Error:\n" +
                               f"{response.info}")
        msg = response.msg  # type: BaseMessage

        condistions_dict = {
            f"condition {i}":
            cdt.replace("<", "").replace(">", "").strip().strip('\n')
            for i, cdt in re.findall(
                r"condition (\d+):\s*(.+?)(?=condition \d+|- Entity)",
                msg.content, re.DOTALL)
        }

        labels = [
            label.strip().strip('\n').strip("\"\'") for label in re.findall(
                r"Entity/Label Recognition of Conditions:\n\[(.+?)\]",
                msg.content, re.DOTALL)[0].split(",")
        ]

        quality = [
            q.replace("<", "").replace(">", "").strip().strip('\n')
            for q in re.findall(
                r"Quality Assessment \(\$Q\$\) \(do not use symbols\):"
                r"\n(.+?)- Iterative", msg.content, re.DOTALL)
        ][0]

        # Convert them into JSON format
        conditions_and_quality_json: Dict[str, Any] = {}
        conditions_and_quality_json["conditions"] = condistions_dict
        conditions_and_quality_json["labels"] = labels
        conditions_and_quality_json["quality"] = quality

        return conditions_and_quality_json
