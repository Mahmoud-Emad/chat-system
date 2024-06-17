class ChatsController < ActionController::API
  before_action :set_application

  def index
    chats = @application.chats
    render json: chats.as_json(only: [ :number, :created_at, :updated_at, :messages_count ])
  end

  def create
    render json: {"message": "This endpoint is only supported on the Go server running on port 8001."}
  end

  def show
    print(params[:token])
    chat = Chat.find_by(application: @application, number: params[:number])

    if chat
      render json: chat.as_json(
        only: [ :number, :created_at, :updated_at, :messages_count ]
      )
    else
      render json: { error: "Chat not found" }, status: :not_found
    end
  end

  private
  def set_application
    @application = Application.find_by(token: params[:application_token])
    render json: { error: "Application is not found" }, status: :not_found unless @application
  end
end
