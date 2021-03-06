# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: field.proto

from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='field.proto',
  package='field',
  syntax='proto3',
  serialized_options=None,
  serialized_pb=b'\n\x0b\x66ield.proto\x12\x05\x66ield\"0\n\x0c\x46ieldRequest\x12\x0c\n\x04slow\x18\x01 \x01(\x08\x12\x12\n\nunreliable\x18\x02 \x01(\x08\"\x1b\n\nFieldReply\x12\r\n\x05\x66ield\x18\x01 \x01(\t2=\n\x05\x46ield\x12\x34\n\x08GetField\x12\x13.field.FieldRequest\x1a\x11.field.FieldReply\"\x00\x62\x06proto3'
)




_FIELDREQUEST = _descriptor.Descriptor(
  name='FieldRequest',
  full_name='field.FieldRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='slow', full_name='field.FieldRequest.slow', index=0,
      number=1, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='unreliable', full_name='field.FieldRequest.unreliable', index=1,
      number=2, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=22,
  serialized_end=70,
)


_FIELDREPLY = _descriptor.Descriptor(
  name='FieldReply',
  full_name='field.FieldReply',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='field', full_name='field.FieldReply.field', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=72,
  serialized_end=99,
)

DESCRIPTOR.message_types_by_name['FieldRequest'] = _FIELDREQUEST
DESCRIPTOR.message_types_by_name['FieldReply'] = _FIELDREPLY
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

FieldRequest = _reflection.GeneratedProtocolMessageType('FieldRequest', (_message.Message,), {
  'DESCRIPTOR' : _FIELDREQUEST,
  '__module__' : 'field_pb2'
  # @@protoc_insertion_point(class_scope:field.FieldRequest)
  })
_sym_db.RegisterMessage(FieldRequest)

FieldReply = _reflection.GeneratedProtocolMessageType('FieldReply', (_message.Message,), {
  'DESCRIPTOR' : _FIELDREPLY,
  '__module__' : 'field_pb2'
  # @@protoc_insertion_point(class_scope:field.FieldReply)
  })
_sym_db.RegisterMessage(FieldReply)



_FIELD = _descriptor.ServiceDescriptor(
  name='Field',
  full_name='field.Field',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=101,
  serialized_end=162,
  methods=[
  _descriptor.MethodDescriptor(
    name='GetField',
    full_name='field.Field.GetField',
    index=0,
    containing_service=None,
    input_type=_FIELDREQUEST,
    output_type=_FIELDREPLY,
    serialized_options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_FIELD)

DESCRIPTOR.services_by_name['Field'] = _FIELD

# @@protoc_insertion_point(module_scope)
