import discord
from datetime import datetime
from sqlalchemy.orm.session import Session

from boarbot.common.botmodule import BotModule
from boarbot.common.config import DATABASE
from boarbot.common.events import EventType
from boarbot.db.kv import KV
from boarbot.db.acl import check_acl_user

debug_users = DATABASE['debugUsers']
DEBUG_USERS_ACL = "boarbot.modules.dbdebug::users"
class DBDebugModule(BotModule):
    async def handle_event(self, db_session: Session, event_type: EventType, args):
        if event_type != EventType.MESSAGE:
            return
        
        message = args[0] # type: discord.Message
        if not (
            message.clean_content.startswith('!kvget') or 
            message.clean_content.startswith('!kvset') or 
            message.clean_content.startswith('!kvdel')):
            return

        if not check_acl_user(db_session, DEBUG_USERS_ACL, message.author.id):
            await self.client.send_message(message.channel, "You are not authorized to use DB debug functionality. I'm telling on you to Santa, too.")
            return
        
        parts = message.clean_content.split(maxsplit=2)
        if parts[0] == '!kvget':
            await self.kvget(db_session, parts, message)
        elif parts[0] == '!kvset':
            await self.kvset(db_session, parts, message)
        elif parts[0] == '!kvdel':
            await self.kvdel(db_session, parts, message)
        else:
            await self.client.send_message(message.channel, f"Unknown DB debug command: {parts[0]}")

    async def kvget(self, db_session: Session, message_parts: [str], message: discord.Message):
        if len(message_parts) < 2:
            await self.client.send_message(message.channel, f"Bad kvget: key is required")
            return
        key = message_parts[1]
        kv = db_session.query(KV).filter(KV.key==key).one_or_none()
        if not kv:
            await self.client.send_message(message.channel, f"kv[`{key}`] -> (missing)")
        else:
            await self.client.send_message(message.channel, f"kv[`{kv.key}`] -> `{kv.value}` @ `{kv.timestamp.isoformat()}`")

    async def kvdel(self, db_session: Session, message_parts: [str], message: discord.Message):
        if len(message_parts) < 2:
            await self.client.send_message(message.channel, f"Bad kvdel: key is required")
            return
        key = message_parts[1]
        kv = db_session.query(KV).filter(KV.key == key).one_or_none()
        if kv:
            db_session.delete(kv)
            db_session.commit()
            await self.client.send_message(message.channel, f"kv[`{kv.key}`] deleted")
        else:
            await self.client.send_message(message.channel, f"kv[`{kv.key}`] not found")

    async def kvset(self, db_session: Session, message_parts: [str], message: discord.Message):
        if len(message_parts) < 3:
            await self.client.send_message(message.channel, f"Bad kvset: key and value are required")
            return
        key = message_parts[1]
        value = message_parts[2]
        kv = db_session.query(KV).filter(KV.key == key).one_or_none()
        if kv:
            kv.value = value
            kv.timestamp = datetime.utcnow()
            await self.client.send_message(message.channel, f"kv[`{kv.key}`] set")
        else:
            kv = KV(key=key, value=value, timestamp=datetime.utcnow())
            db_session.add(kv)
            await self.client.send_message(message.channel, f"kv[`{kv.key}`] created")
        db_session.commit()
