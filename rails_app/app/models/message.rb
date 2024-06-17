class Message < ApplicationRecord
  include Elasticsearch::Model
  include Elasticsearch::Model::Callbacks

  belongs_to :chat

  before_create :set_number

  settings index: { number_of_shards: 1 } do
    mappings dynamic: "false" do
      indexes :body, type: "text"
      indexes :chat_number, type: "integer"
    end
  end

  def as_indexed_json(options = {})
    self.as_json(only: [ :body, :chat_number ])
  end

  private

  def set_number
    self.number = (chat.messages.maximum(:number) || 0) + 1
  end
end
