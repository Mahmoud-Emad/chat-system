class MessagesController < ActionController::API
  before_action :set_application, :set_chat

  def index
    messages = @chat.messages
    render json: messages.as_json(
      only: [ :number, :created_at, :updated_at, :body ]
    )
  end

  def show
    messages = Message.find_by(chat: @chat, number: params[:number])

    if messages
      render json: messages.as_json(
        only: [ :number, :created_at, :updated_at, :body ]
      )
    else
      render json: { error: "Message not found" }, status: :not_found
    end
  end

  def create
    if params["body"].blank?
      render json: { error: "The message body is required, Try to send `{ body: 'Simple Body!' }` as the request body." }, status: 400
      return
    end

    message = @chat.messages.build(message_params)
    if message.save
      render json: message.as_json(
        only: [ :number, :created_at, :updated_at, :body ]
      ), status: :created
    else
      render json: message.errors, status: :unprocessable_entity
    end
  end

  def search
    if params[:query].blank?
      render json: { error: "Query parameter is required for search" }, status: :bad_request
      return
    end

    messages = Message.search({
      query: {
        bool: {
          must: {
            match: {
              body: params[:query]
            }
          },
          filter: {
            term: { chat_number: @chat.number }
          }
        }
      }
    })

    response = messages.results
    render json: response
  end

  private

  def set_application
    @application = Application.find_by(token: params[:application_token])
    render json: { error: "Application is not found" }, status: :not_found unless @application
  end

  def set_chat
    @chat = Chat.find_by(application: @application, number: params[:chat_number])
    render json: { error: "Chat is not found" }, status: :not_found unless @chat
  end

  def message_params
    params.require(:message).permit(:body)
  end
end
