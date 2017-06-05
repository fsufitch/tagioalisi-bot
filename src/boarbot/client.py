import discord

from boarbot.dispatch import dispatch_event
from boarbot.common.events import EventType

class BoarBotClient(discord.Client):
    async def on_ready(self):
        await dispatch_event(EventType.READY)

    async def on_reload(self):
        await dispatch_event(EventType.RELOAD)

    async def on_message(self, message: discord.Message):
        await dispatch_event(EventType.MESSAGE, message)

    async def on_messsage_delete(self, message: discord.Message):
        await dispatch_event(EventType.MESSAGE_DELETE, message)

    async def on_message_edit(self, before: discord.Message, after: discord.Message):
        await dispatch_event(EventType.MESSAGE_EDIT, before, after)

    async def on_reaction_add(self, reaction: discord.Reaction, user: discord.User):
        await dispatch_event(EventType.REACTION_ADD, reaction, user)

    async def on_reaction_remove(self, reaction: discord.Reaction, user: discord.User):
        await dispatch_event(EventType.REACTION_REMOVE, reaction, user)

    async def on_reaction_clear(self, message: discord.Message, reactions: [discord.Reaction]):
        await dispatch_event(EventType.REACTION_CLEAR, message, reactions)

    async def on_member_join(self, member: discord.Member):
        await dispatch_event(EventType.MEMBER_JOIN, member)

    async def on_member_remove(self, member: discord.Member):
        await dispatch_event(EventType.MEMBER_REMOVE, member)

    async def on_member_update(self, before: discord.Member, after: discord.Member):
        await dispatch_event(EventType.MEMBER_UPDATE, before, after)
