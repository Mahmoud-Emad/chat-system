class Chat < ApplicationRecord
  belongs_to :application
  has_many :messages

  before_create :set_number

  private

  def set_number
    self.number = (application.chats.maximum(:number) || 0) + 1
  end
end
