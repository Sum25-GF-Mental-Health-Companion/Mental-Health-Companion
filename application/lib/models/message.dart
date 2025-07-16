class Message {
  final int? sessionId;
  final String sender;
  final String content;

  Message({
    required this.sender,
    required this.content,
    this.sessionId,
  });

  Map<String, dynamic> toJson() {
    return {
      'session_id': sessionId,
      'sender': sender,
      'content': content,
    };
  }

  factory Message.fromJson(Map<String, dynamic> json) {
    return Message(
      sessionId: json['session_id'],
      sender: json['sender'],
      content: json['content'],
    );
  }
}
