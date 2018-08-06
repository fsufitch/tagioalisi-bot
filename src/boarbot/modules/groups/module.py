import discord

from boarbot.common.botmodule import BotModule
from boarbot.common.config import CONFIG
from boarbot.common.events import EventType
from boarbot.common.chunks import chunk_lines

from .cmd import GROUPS_PARSER, GroupsParserException

GROUP_MANAGER_IDS = CONFIG.get('groups', {}).get('managers', [])
GROUP_PREFIX = 'g-'
GROUPS_COMMAND = '!groups'
ERROR_FORMAT = '`{error}`\nTry `!groups --help` to get usage instructions.'

class GroupsModule(BotModule):
    async def handle_event(self, event_type: EventType, args):
        if event_type != EventType.MESSAGE:
            return
        message = args[0]

        if message.author.bot:
            return # Ignore bots

        args = self.parse_command(GROUPS_COMMAND, message)
        if args is None:
            return

        if not message.server:
            await self.client.send_message(message.channel, "Please send the message in a server-context channel, not a direct message")
            return

        try:
            parsed_args = GROUPS_PARSER.parse_args(args)
        except GroupsParserException as e:
            await self.client.send_message(message.channel, ERROR_FORMAT.format(error=e.args[0]))
            return

        if parsed_args.help:
            await self.client.send_message(message.channel, '```' + GROUPS_PARSER.format_help() + '```')
            return

        command_map = {
            'list': self.list_groups,
            'members': self.list_members,
            'create': self.create_group,
            'delete': self.delete_group,
            'join': self.join_group,
            'leave': self.leave_group,
        }
        command = parsed_args.command
        if command not in command_map:
            await self.client.send_message(message.channel, ERROR_FORMAT.format(error="no such command: %s" % command))
            return

        group_name = parsed_args.group
        mentions = self._get_server_members(message, parsed_args.users)

        await command_map[command](message, group_name, mentions)

    async def list_groups(self, trigger_message: discord.Message, *args):
        groups = self._get_group_list(trigger_message)
        lines = []
        for group in groups:
            group_name = group.name[len(GROUP_PREFIX):]
            line = '- `{name}` - mention with: `@{prefix}{name}` or `{mention}`'.format(name=group_name, prefix=GROUP_PREFIX, mention=group.mention)
            lines.append(line)

        if lines:
            await self.client.send_message(trigger_message.author, '\n'.join(lines))
        else:
            await self.client.send_message(trigger_message.channel, 'No mentionable groups with prefix `%s` found' % GROUP_PREFIX)
        await self.client.send_message(trigger_message.channel, 'Group list sent via private message.')

    async def list_members(self, trigger_message: discord.Message, group_name: str, *args):
        role_name = GROUP_PREFIX+group_name
        group = self._find_group(trigger_message, group_name)
        if not group:
            await self.client.send_message(trigger_message.channel, 'ERROR: group with name `%s` not found (role `%s`)' % (group_name, role_name))
            return

        group_members = []
        for member in trigger_message.server.members:
            if group in member.roles:
                group_members.append(member)

        output_lines = ['Group with name %s (role `%s`) has %d members.' % (group_name, role_name, len(group_members))]
        output_lines += ['- %s' % (member.nick or member.name) for member in group_members]
        for message_chunk in chunk_lines(output_lines):
            reply = '\n'.join(message_chunk)
            await self.client.send_message(trigger_message.author, reply)
        await self.client.send_message(trigger_message.channel, 'Member list for group `%s` (role `%s`) sent via private message.' % (group_name, role_name))

    async def create_group(self, trigger_message: discord.Message, group_name: str, *args):
        if not self._is_group_manager(trigger_message.author):
            await self.client.send_message(trigger_message.channel, 'ERROR: %s is not a server manager, and cannot create groups' % trigger_message.author.mention)
            return

        role_name = GROUP_PREFIX+group_name
        groups = self._get_group_list(trigger_message)
        if any(g.name==group_name for g in groups):
            await self.client.send_message(trigger_message.channel, 'ERROR: group with name `%s` already exists (role `%s`)' % (group_name, role_name))
            return

        await self.client.create_role(trigger_message.server, name=role_name, hoist=False, mentionable=True)
        await self.client.send_message(trigger_message.channel, 'Created group with name `%s` (role `%s`)' % (group_name, role_name))


    async def delete_group(self, trigger_message: discord.Message, group_name: str, *args):
        if not self._is_group_manager(trigger_message.author):
            await self.client.send_message(trigger_message.channel, 'ERROR: %s is not a server manager, and cannot delete groups' % trigger_message.author.mention)
            return

        role_name = GROUP_PREFIX+group_name
        group = self._find_group(trigger_message, group_name)
        if not group:
            await self.client.send_message(trigger_message.channel, 'ERROR: group with name `%s` not found (role `%s`)' % (group_name, role_name))
            return
        await self.client.delete_role(trigger_message.server, group)
        await self.client.send_message(trigger_message.channel, 'Deleted group with name `%s` (role `%s`)' % (group_name, role_name))

    async def join_group(self, trigger_message: discord.Message, group_name: str, user_mentions: [discord.Member]):
        if user_mentions:
            if not self._is_group_manager(trigger_message.author):
                await self.client.send_message(trigger_message.channel, 'ERROR: %s is not a server manager, and cannot use user mentions to edit groups' % trigger_message.author.mention)
                return
        else:
            user_mentions = [trigger_message.author]

        role_name = GROUP_PREFIX+group_name
        group = self._find_group(trigger_message, group_name)
        if not group:
            await self.client.send_message(trigger_message.channel, 'ERROR: group with name `%s` not found (role `%s`)' % (group_name, role_name))
            return

        for user in user_mentions:
            await self.client.add_roles(user, group)

        added_users = ', '.join(u.mention for u in user_mentions)
        message = 'Added {users} to group `{group}` (role `{role}`)'.format(users=added_users, group=group_name, role=role_name)
        await self.client.send_message(trigger_message.channel, message)

    async def leave_group(self, trigger_message: discord.Message, group_name: str, user_mentions: [discord.Member]):
        if user_mentions:
            if not self._is_group_manager(trigger_message.author):
                await self.client.send_message(trigger_message.channel, 'ERROR: %s is not a server manager, and cannot use user mentions to edit groups' % trigger_message.author.mention)
                return
        else:
            user_mentions = [trigger_message.author]

        role_name = GROUP_PREFIX+group_name
        group = self._find_group(trigger_message, group_name)
        if not group:
            await self.client.send_message(trigger_message.channel, 'ERROR: group with name `%s` not found (role `%s`)' % (group_name, role_name))
            return

        for user in user_mentions:
            await self.client.remove_roles(user, group)

        removed_users = ', '.join(u.mention for u in user_mentions)
        message = 'Removed {users} from group `{group}` (role `{role}`)'.format(users=removed_users, group=group_name, role=role_name)
        await self.client.send_message(trigger_message.channel, message)


    def _get_group_list(self, trigger_message: discord.Message) -> [discord.Role]:
        groups = [] # type: [discord.Role]
        for role in trigger_message.server.roles:
            if (role.name.startswith(GROUP_PREFIX) and role.mentionable):
                groups.append(role)
        return groups

    def _find_group(self, trigger_message: discord.Message, group_name: str) -> discord.Role:
        role_name = GROUP_PREFIX+group_name
        groups = self._get_group_list(trigger_message)
        for group in groups:
            if group.name == role_name:
                return group
        return None

    def _get_server_members(self, trigger_message: discord.Message, mentions: [str]) -> [discord.Member]:
        members = [] # type: [discord.Member]
        for member in trigger_message.server.members:
            if member.mention in mentions:
                members.append(member)
        return members

    def _is_group_manager(self, user: discord.User) -> bool:
        return user.id in GROUP_MANAGER_IDS
