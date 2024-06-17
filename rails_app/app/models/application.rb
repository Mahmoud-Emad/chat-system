class Application < ApplicationRecord
  has_many :chats
  after_commit :flush_cache
  before_create :generate_token

  def self.cached_all
    Rails.cache.fetch("applications/all", expires_in: 1.hour) do
      all.to_a
    end
  end

  private

  def flush_cache
    Rails.cache.delete("applications/all")
  end

  def generate_token
    self.token = SecureRandom.hex(10)
  end
end
